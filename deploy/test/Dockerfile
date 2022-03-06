FROM golang:1.13

WORKDIR /go/src/github.com/sifchain/al-jabaal

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make install

# Initialize the genesis.json file that will help you to bootstrap the network
RUN ebd init local --chain-id=al-jabaal

# Configure your cli to use the keyring-backend test so that you don't need to enter a password
# **_NOTE_** Don't use the test backend on production
RUN ebcli config keyring-backend test

# Add some more configuration to avoid using flags
RUN ebcli config chain-id al-jabaal
RUN ebcli config trust-node true
RUN ebcli config indent true
RUN ebcli config output json

# Create a key to hold your validator account and for another test account
RUN ebcli keys add validator
RUN ebcli keys add testuser

# Initialize the genesis account and transaction
RUN ebd add-genesis-account $(ebcli keys show validator -a) 1000000000stake,1000000000atom

# Create genesis transaction
RUN ebd gentx --name validator --keyring-backend test

# Collect genesis transaction
RUN ebd collect-gentxs
