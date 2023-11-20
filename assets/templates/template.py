import sys
import os
import requests
import pwn
import re
import json
import base64
import string
from tqdm import tqdm

url = sys.argv[1]
port = ...
headers = {
    'User-Agent':'checker'
}

#===========================================================
#                    EXPLOIT GOES HERE
#===========================================================
def main():
    req = requests.get(url=url, port=port, headers=headers)
    resp = req.response()

#===========================================================
if __name__ == '__main__':
    main()