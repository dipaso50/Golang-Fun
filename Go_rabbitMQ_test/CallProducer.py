import requests
import json, sys

repeticiones = int(sys.argv[1])
count = 0
while (count < repeticiones):    
    msg = "Message " + str(count)
    count = count + 1
    print("Calling producer Msg--> " + msg)
    r = requests.post("http://localhost:3000/", json.dumps({'Data': msg})) 

 