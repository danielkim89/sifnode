import logging
import os
import time
import json
import pytest
import string
import random

import burn_lock_functions
from burn_lock_functions import EthereumToSifchainTransferRequest
import test_utilities
from pytest_utilities import generate_test_account
from integration_env_credentials import sifchain_cli_credentials_for_test
from test_utilities import get_required_env_var, SifchaincliCredentials, get_optional_env_var, ganache_owner_account, \
    get_shell_output_json, get_shell_output, detect_errors_in_sifnodecli_output, get_transaction_result, amount_in_wei

# CODE TO GENERATE RANDOM STRING FOR DISPENSATION NAME AS DISPENSATION NAME IS A UNIQUE KEY
def id_generator(size=6, chars=string.ascii_uppercase + string.digits):
    return ''.join(random.choice(chars) for _ in range(size))

# CODE TO GENERATE NEW ADDRESS AND ADD IT IN THE KEYRING
def create_new_sifaddr_and_key():
    new_account_key = test_utilities.get_shell_output("uuidgen").replace("-", "")
    credentials = sifchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_sifaddr(credentials=credentials, keyname=new_account_key)
    return new_addr["address"], new_addr["name"]


# CODE TO SEND SOME SAMPLE TOKEN TO NEW ADDRESS
def send_sample_rowan(from_address, to_address, amount, keyring_backend, chain_id):
    logging.debug(f"transfer_rowan")
    sifchain_fees_entry = f"--fees 150000rowan"
    keyring_backend_entry = f"--keyring-backend {keyring_backend}"
    output = 'output.json'
    cmd = " ".join([
        "sifnodecli tx send",
        f"{from_address}",
        f"{to_address}",
        f"{amount}",
        keyring_backend_entry,
        sifchain_fees_entry,
        f"--chain-id {chain_id}",
        f"--yes -o json"

    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    return json_str


# CODE TO QUERY BLOCK FOR NEW DISPENSATION TXN
def query_block_claim(txnhash):
    cmd = " ".join([
        "sifnodecli q tx",
        f"{txnhash}",
        "--chain-id localnet",
        "-o json"
    ])
    json_str = get_shell_output_json(cmd)
    return json_str


# CODE TO CHECK ACCOUNT BALANCE
def balance_check(address, currency):
    logging.debug(f"check_balance")
    cmd = " ".join([
        "sifnodecli query account",
        f"{address}",
        f"-o json"
    ])
    json_str = get_shell_output_json(cmd)
    amountbalance = json_str['value']['coins']
    for i in amountbalance:
        if i['denom'] == currency:
            balance = i['amount']
    return (balance)

#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY ONLINE DISPENSATION TXN
def create_online_singlekey_txn(
        claimType,
        signing_address,
        chain_id,
        sifnodecli_node
):
    logging.debug(f"create_online_dispensation")
    sifchain_fees_entry = f"--gas auto"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "sifnodecli tx dispensation create",
        f"{claimType}",
        output,
        sifchain_fees_entry,
        f"--fees 50000rowan",
        f"--from {signing_address}",
        f"--chain-id={chain_id}",
        f"{sifnodecli_node}",
        keyring_backend_entry,
        f"--yes -o json"

    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn

#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY ONLINE DISPENSATION TXN WITH AN ASYNC FLAG
def create_online_singlekey_async_txn(
        claimType,
        signing_address,
        chain_id,
        sifnodecli_node
):
    logging.debug(f"create_online_dispensation")
    sifchain_fees_entry = f"--gas auto"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "sifnodecli tx dispensation create",
        f"{claimType}",
        output,
        sifchain_fees_entry,
        f"--fees 50000rowan",
        f"--from {signing_address}",
        f"--chain-id={chain_id}",
        f"{sifnodecli_node}",
        keyring_backend_entry,
        f"--broadcast-mode async",
        f"--yes -o json"

    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn

#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY OFFLINE DISPENSATION TXN
def create_offline_singlekey_txn(
        claimType,
        signing_address,
        chain_id,
        sifnodecli_node
    ):
    logging.debug(f"create_unsigned_offline_dispensation_txn")
    sifchain_fees_entry = f"--gas auto"
    output = 'output.json'
    cmd = " ".join([
        "sifnodecli tx dispensation create",
        f"{claimType}",
        output,
        f"--from {signing_address}",
        f"--chain-id={chain_id}",
        f"{sifnodecli_node}",
        f"--fees 150000rowan",
        f"--generate-only", 
        f"--yes -o json"
        
    ])
    json_str = get_shell_output_json(cmd)
    assert(json_str.get("code", 0) == 0)
    return json_str

#CODE TO SIGN DISPENSATION TXN BY A USER
def sign_txn(signingaddress, offlinetx):
    keyring_backend_entry = f"--keyring-backend test"
    cmd = " ".join([
        "sifnodecli tx sign",
        f"--from {signingaddress}",
        f"{offlinetx}",
        keyring_backend_entry,
        "--chain-id localnet",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    return json_str

#CODE TO BROADCAST SINGLE SIGNED TXN ON BLOCK
def broadcast_txn(signedtx):
    cmd = " ".join([
        "sifnodecli tx broadcast",
        f"{signedtx}",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    txn = json_str["txhash"]
    return txn

#CODE TO BROADCAST SINGLE SIGNED TXN ON BLOCK WITH AN ASYNC FLAG
def broadcast_async_txn(signedtx):
    cmd = " ".join([
        "sifnodecli tx broadcast",
        f"{signedtx}",
        f"--broadcast-mode async",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    txn = json_str["txhash"]
    return txn