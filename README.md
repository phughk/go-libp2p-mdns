#### Library for Peer Discovery using MDNS

<b>"This Library is still work in progress!!!"</b>

The Library follows [Libp2p MDNS specification](https://github.com/libp2p/specs/blob/master/discovery/mdns.md) for peer discovery. 


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

<i>This project is a by product of the [bounty](https://github.com/ethberlinzwei/Bounties/issues/19) in the [Eth Hack Zwei](https://github.com/ethberlinzwei/KnowledgeBase). Special thanks to @raulk </i>
