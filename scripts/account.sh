#!/bin/sh

mnemonic="april patch recipe debate remove hurdle concert gesture design near predict enough color tail business imitate twelve february punch cheap vanish december cool wheel"
mnemonic1="soft brown armed regret hip few ancient control steel bright basic swamp sentence present immune napkin orbit giggle year another crowd essence noble dice"
mnemonic2="bomb sand fashion torch return coconut color captain vapor inhale lyrics lady grant ordinary lazy decrease quit devote paddle impulse prize equip hip ball"
mnemonic3="road gallery tooth script volcano deputy summer acid bulk anger fatigue notable secret blood bean apology burger rookie rug bench away dutch secret upper"
echo "$mnemonic" | gstaked keys add oracle1 --recover --keyring-backend=test --home /gstaked
echo "$mnemonic1" | gstaked keys add oracle2 --recover --keyring-backend=test --home /gstaked
echo "$mnemonic2" | gstaked keys add oracle3 --recover --keyring-backend=test --home /gstaked
echo "$mnemonic3" | gstaked keys add oracle4 --recover --keyring-backend=test --home /gstaked