import requests

import requests


params = {
    "coilId": "H19148841L",
    "curDir": "E:/2250hrm/201911/20191108",
    "factorName": "wedge40",

    "functionName": "aimRate",
    "aim": 0,
    "tolerance": 20,
    "unit": "um",

    "lengthName": "main",
    "headLen": 17,
    "tailLen": 27,
    "headCut": 5,
    "tailCut": 5,

}
r = requests.get('http://192.168.88.158:8999/api/v1/ponds/stats/H19148841L',
                 params=params)

print(r.url)
print(r.json())
