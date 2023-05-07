#!/bin/sh

rm -rf ~/.gstaked/
rm -rf ~/.gstaked/

mnemonic="together chief must vocal account off apart dinosaur move canvas spring whisper improve cruise idea earn reflect flash goat illegal mistake blood earn ridge"
mnemonic1="marble allow december print trial know resource cry next segment twice nose because steel omit confirm hair extend shrimp seminar one minor phone deputy"
mnemonic2="axis decline final suggest denial erupt satisfy weekend utility fortune dry glory recall real other evil spatial speed seek rubber struggle wolf tortoise large"
mnemonic3="knock board dolphin cricket strike sense throw security mistake link ocean educate merit pet public economy embark shoot horror pond budget rent toe frozen"

gstaked init test --chain-id test

echo "$mnemonic" | gstaked keys add test --recover --keyring-backend=test

echo "$mnemonic1" | gstaked keys add test1 --recover --keyring-backend=test

echo "$mnemonic2" | gstaked keys add test2 --recover --keyring-backend=test

echo "$mnemonic3" | gstaked keys add test3 --recover --keyring-backend=test

gstaked add-genesis-account test 10000000000000000000stake,1000000000ustkstake  --keyring-backend=test
gstaked gentx test 100000000stake --chain-id test --keyring-backend=test
gstaked collect-gentxs


#gstaked start


