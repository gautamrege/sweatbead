## The Sweat Bead ##

This is an entirely hypothetical attempt to build a platform using differnet microservices. 
In this example, we don't concentrate on the hardware but only the process of building different 
microservices incorporating the Single Responsibilty principle.

Did you know that Sweat can be used for various health deductions including diabeties? Well, I have
still to find a really good sweat analyzer but we shall assume that a sweat-bead is one such device.

### The Architecture ###

We manage multiple microservices microservices to manage sweat and user/device data, to provide
alerts and reports. 


The SweatMgr manages all Sweat data.
The UserMgr manages all user data
The ReportsMgr is a client to query the SweatMgr

This is a Hands-on tutorial and you're expected to make it work (not see it work)

