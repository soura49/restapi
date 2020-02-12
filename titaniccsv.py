#Python Function to insert all the data from the csv file into the DB using the curl command

import csv
import json
import os

with open('titanic.csv') as csv_file:
    csv_reader = csv.reader(csv_file, delimiter=',')
    line_count = 0
    output = {}
    for row in csv_reader:
        output["age"] = row[4]
        output["parentsOrChildrenAboard"] = row[6]
        output["siblingsOrSpousesAboard"] = row[7]
        output["fare"] = row[7]
        output["sex"] = row[3]
        output["survived"] = row[0]
        output["passengerClass"] = row[1]
        output["name"] = row[2]
        ou = json.dumps(output)
        cmd = "curl -X POST -d " +"\'"+ou+"\'"+ " http://127.0.0.1:8088/people"
        os.system(cmd)