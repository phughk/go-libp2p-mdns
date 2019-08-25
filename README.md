#### Library for Peer Discovery using MDNS

<b>"This Library is still work in progress!!!"</b>

The Library follows [Libp2p MDNS specification](https://github.com/libp2p/specs/blob/master/discovery/mdns.md) for peer discovery. 

##### Directory Structure

.
├── ...
├── mdns.go                      # Initialize the Server
├── marshal.go                   # Serialize the mdns packets
├── unmarshal.go                 # UnSerialize mdns packets
├── service.go                   # The Services exposed my Mdns (eg: Poll(), ReadPacket(), etc)
├── handler.go                   # Event handler Interface Peer Discovery and Network Interrupt
├── log.go                       # Log Managment
└── ...

##### Implementation Status

Parts that are Implemented 

*   Marshalar and Unmarshalar for the Querry (BuildQuery(), UnpackMessage())
*   Initialization of Reuse Port (Init())
*   Handler Interface
*   Send Querry

Parts that are partially Implemented

*   Marshalar and Unmarshalar for the Response Message
*   ParsePacket and Call the Handler

Parts that Need to be Implemented 

*   UnitTests
*   Examples
*   Code Documentation
*   Better Error Handling

