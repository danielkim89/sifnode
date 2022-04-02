import time

from siftool import eth, test_utils, sifchain
from siftool.common import *
import web3
from json_file_writer import write_registry_json, ibc_json_file_path, double_peggy_json_file_path

fund_amount_eth = eth.ETH
rowan_unit = 10 ** 18
fund_amount_sif = 2 * rowan_unit
rowan_contract_address = web3.Web3.toChecksumAddress('0x5fbdb2315678afecb367f032d93f642f64180aa3')
ibc_token_symbol = 'ibc/FEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACEFEEDFACE'
double_peggy_symbol = 'sifBridge00030x1111111111111111111111111111111111111111'
double_peggy_address = '0x1111111111111111111111111111111111111111'

fund_amount_ibc = 10000
fund_amount_double_peggy = 10000

admin_account = "sifnodeadmin"

def test_rowan_to_eth_and_back_to_sifnode_transfer_valid(ctx):
    # get rowan contract
    rowan_sc = ctx.get_generic_erc20_sc(rowan_contract_address)

    # Create/retrieve a test ethereum account
    test_eth_account = ctx.create_and_fund_eth_account(fund_amount=fund_amount_eth)

    # create/retrieve a test sifchain account
    test_sif_account = ctx.create_sifchain_addr(fund_amounts=[[fund_amount_sif, "rowan"], [fund_amount_eth, ctx.ceth_symbol]])

    # init balance for erc20 rowan
    test_eth_account_initial_balance = ctx.get_erc20_token_balance(rowan_contract_address, test_eth_account)

    # sif account initial balance
    test_sif_account_initial_balance = ctx.get_sifchain_balance(test_sif_account)

    # lock rowan in sifnode
    amount_to_lock = 1 * rowan_unit
    ctx.sifnode_client.send_from_sifchain_to_ethereum(test_sif_account, test_eth_account, amount_to_lock, "rowan")
    ctx.advance_blocks()
    test_sif_account_after_lock_balance = ctx.wait_for_sif_balance_change(test_sif_account, test_sif_account_initial_balance, [[amount_to_lock, "rowan"]])

    # we need take the transaction fee into account
    rowan_balance = test_sif_account_initial_balance["rowan"] - amount_to_lock
    assert rowan_balance >= test_sif_account_after_lock_balance["rowan"]

    # wait the ethereum reciever's rowan balance change
    eth_account_final_balance = ctx.wait_for_eth_balance_change(test_eth_account, test_eth_account_initial_balance, token_addr=rowan_contract_address, timeout=90)

    # check the rowan balance as expected
    assert eth_account_final_balance == test_eth_account_initial_balance + amount_to_lock

    test_sif_account_before_receive = ctx.get_sifchain_balance(test_sif_account)

    ctx.send_from_ethereum_to_sifchain(test_eth_account, test_sif_account, amount_to_lock, token_sc=rowan_sc, isLock=False)
    ctx.advance_blocks()

    test_sif_account_after_receive = ctx.wait_for_sif_balance_change(test_sif_account, test_sif_account_before_receive, [[amount_to_lock, "rowan"]])

    assert test_sif_account_after_receive["rowan"] == amount_to_lock + test_sif_account_before_receive["rowan"]

def test_double_peggy_token(ctx):
    # Create/retrieve a test ethereum account
    test_eth_account = ctx.create_and_fund_eth_account(fund_amount=fund_amount_eth)

    # create/retrieve a test sifchain account
    test_sif_account = ctx.create_sifchain_addr(
        fund_amounts=[[fund_amount_sif, "rowan"], [fund_amount_eth, ctx.ceth_symbol],
                      [fund_amount_double_peggy, double_peggy_symbol]])

    destination_contract_address = ctx.get_destination_contract_address(double_peggy_address)

    if destination_contract_address == eth.NULL_ADDRESS:
        # write the ibc token entry to a json file
        write_registry_json(double_peggy_symbol, double_peggy_address, isIbc=False)

        # register the ibc token
        ctx.sifnode.peggy2_token_registry_register_all(double_peggy_json_file_path, [0.5, "rowan"], 1.5, admin_account,
                                                       ctx.sifnode_chain_id)
        init_test_eth_account_balance = 0
    else:
        init_test_eth_account_balance = ctx.get_erc20_token_balance(destination_contract_address, test_eth_account)

    ctx.sifnode_client.send_from_sifchain_to_ethereum(test_sif_account, test_eth_account, fund_amount_double_peggy, double_peggy_symbol)
    ctx.advance_blocks()

    if destination_contract_address == eth.NULL_ADDRESS:
        created_contract_address = ctx.wait_for_new_bridge_token_created(double_peggy_address)
        double_peggy_token_sc = ctx.get_generic_erc20_sc(created_contract_address)
        final_test_eth_account_balance = ctx.get_erc20_token_balance(created_contract_address, test_eth_account)
    else:
        double_peggy_token_sc = ctx.get_generic_erc20_sc(destination_contract_address)
        final_test_eth_account_balance = ctx.wait_for_eth_balance_change(test_eth_account, init_test_eth_account_balance, token_addr=destination_contract_address)

    assert init_test_eth_account_balance + fund_amount_double_peggy == final_test_eth_account_balance

    init_test_sif_account_balance = ctx.get_sifchain_balance(test_sif_account)

    ctx.send_from_ethereum_to_sifchain(test_eth_account, test_sif_account, fund_amount_double_peggy, token_sc=double_peggy_token_sc, isLock=False)
    final_test_sif_account_balance = ctx.wait_for_sif_balance_change(test_sif_account, init_test_sif_account_balance, [[fund_amount_double_peggy, double_peggy_symbol]])
    assert sifchain.balance_delta(init_test_sif_account_balance, final_test_sif_account_balance)[double_peggy_symbol] == fund_amount_double_peggy

def test_ibc_to_eth_and_back_to_sifnode_transfer_valid(ctx):
    # deploy erc20 for ibc token
    token_data: test_utils.ERC20TokenData = ctx.generate_random_erc20_token_data()
    ibc_token_sc = ctx.deploy_new_generic_erc20_token(token_data.name, token_data.symbol, 18, cosmosDenom=ibc_token_symbol)

    # add bridge token into whitelist
    ctx.bridge_bank_add_existing_bridge_token(ibc_token_sc.address)

    # ctx.tx_grant_minter_role(ibc_token_sc, ctx.get_cosmos_bridge_sc().address)
    ctx.tx_grant_minter_role(ibc_token_sc, ctx.get_bridge_bank_sc().address)

    # Create/retrieve a test ethereum account
    test_eth_account = ctx.create_and_fund_eth_account(fund_amount=fund_amount_eth)

    # create/retrieve a test sifchain account
    test_sif_account = ctx.create_sifchain_addr(
        fund_amounts=[[fund_amount_sif, "rowan"], [fund_amount_eth, ctx.ceth_symbol], [fund_amount_ibc, ibc_token_symbol]])

    # write the ibc token entry to a json file
    write_registry_json(ibc_token_symbol, ibc_token_sc.address, isIbc=True)

    # register the ibc token
    ctx.sifnode.peggy2_token_registry_register_all(ibc_json_file_path, [0.5, "rowan"], 1.5, admin_account, ctx.sifnode_chain_id)

    test_eth_account_init_balance = ctx.get_erc20_token_balance(ibc_token_sc.address, test_eth_account)

    ctx.sifnode_client.send_from_sifchain_to_ethereum(test_sif_account, test_eth_account, fund_amount_ibc, ibc_token_symbol)
    test_eth_account_after_lock = ctx.wait_for_eth_balance_change(test_eth_account, test_eth_account_init_balance, token_addr=ibc_token_sc.address, timeout=90)

    assert test_eth_account_after_lock == test_eth_account_init_balance + fund_amount_ibc

    test_sif_account_init_balance = ctx.get_sifchain_balance(test_sif_account)

    ctx.send_from_ethereum_to_sifchain(test_eth_account, test_sif_account, fund_amount_ibc, token_sc=ibc_token_sc,
                                       isLock=False)
    ctx.advance_blocks()

    test_sif_account_after_receive = ctx.wait_for_sif_balance_change(test_sif_account, test_sif_account_init_balance, [[fund_amount_ibc, ibc_token_symbol]])

    assert sifchain.balance_delta(test_sif_account_init_balance, test_sif_account_after_receive)[ibc_token_symbol] == fund_amount_ibc




