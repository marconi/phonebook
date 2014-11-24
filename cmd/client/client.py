import sys
import uuid
from datetime import datetime
sys.path.append('../../services/py')

from thrift.transport import TSocket, TTransport
from thrift.protocol import TBinaryProtocol
from contact import ContactSvc, ttypes

socket = TSocket.TSocket('localhost', 9090)
transport = TTransport.TFramedTransport(socket)
protocol = TBinaryProtocol.TBinaryProtocol(transport)
client = ContactSvc.Client(protocol)

transport.open()

c1 = ttypes.Contact(uuid.uuid4().hex, 'Bob', '111-1111', 'bob@wonderland.com', datetime.now().isoformat())
c1 = client.create(c1)

contacts = client.fetch()
print contacts
