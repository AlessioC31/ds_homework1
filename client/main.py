import requests

BASEURL = 'http://10.84.146.6:8080'

def getContentLength(encoding, size):
    r = requests.get(f'{BASEURL}/{encoding}/{size}')

    return r.headers['content-length']


if __name__ == '__main__':
    print(f"JSON-encoding, tiny: {getContentLength('json', 'tiny')} bytes")
    print(f"JSON-encoding, medium: {getContentLength('json', 'medium')} bytes")
    print(f"JSON-encoding, large: {getContentLength('json', 'large')} bytes")

    print(f"Protobuf-encoding, tiny: {getContentLength('protobuf', 'tiny')} bytes")
    print(f"Protobuf-encoding, medium: {getContentLength('protobuf', 'medium')} bytes")
    print(f"Protobuf-encoding, large: {getContentLength('protobuf', 'large')} bytes")

    print(f"XML-encoding, tiny: {getContentLength('xml', 'tiny')} bytes")
    print(f"XML-encoding, medium: {getContentLength('xml', 'medium')} bytes")
    print(f"XML-encoding, large: {getContentLength('xml', 'large')} bytes")