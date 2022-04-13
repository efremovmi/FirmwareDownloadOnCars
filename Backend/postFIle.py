import argparse
import requests

parser = argparse.ArgumentParser(description='Piece of shit')

parser.add_argument("--filename", type=str, help="Your filename")
parser.add_argument("--ip", type=str, help="Your ip")


args = parser.parse_args()
filename = args.filename
ip = args.ip

requests.post('http://{}/mycar'.format(ip), files={'firmware': open(filename, 'rb')})