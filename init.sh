#!/bin/bash

rm -r ~/.ebd
rm -r ~/.ebcli

ebd init local --chain-id=al-jabaal

ebcli config keyring-backend test
ebcli config chain-id al-jabaal
ebcli config trust-node true
ebcli config indent true
ebcli config output json

ebcli keys add validator
ebcli keys add testuser
ebd add-genesis-account $(ebcli keys show validator -a) 1000000000stake,1000000000atom

ebd gentx --name validator --keyring-backend test
ebd collect-gentxs