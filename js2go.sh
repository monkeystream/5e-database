MODEL_FILE=./translator/models/models.go

rm $MODEL_FILE
echo "package models" >> $MODEL_FILE

for entry in ./src/*
do
    if [ "${entry##*.}" = "json" ]; then
        json2struct -f $entry | sed -e "s|JSONToStruct|${entry}|g" |        \
            sed -e "s|package generated||g" | sed -e "s|./src/5e-||g" |     \
            sed -e "s|-||g" | sed -e "s|\.json \[\]| |gm" >> $MODEL_FILE
    fi
done 
