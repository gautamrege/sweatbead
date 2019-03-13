## The Sweat Bead ##

This is an entirely hypothetical attempt to build a scalable platform using microservices. 
In this example, we don't concentrate on the hardware but only scaling large number of 
events coming in from pseudo-devices. And what do we measure? *Sweat* :)

### The Architecture ###

We manage multiple services (also called microservices) to manage a large platform of data,
alerts and reports. Let's list down the different services 

#### Profile Manager ####

This microservice is responsible for user management:

* Regitration & Login
* User Profile
* User Dashboard (prifile data for daily and weekly information)

#### Events Manager ####

This microservice is responsible for all event data:

* Receive event data from the sweat-bead.
* Process the event data / calculate for alert
* Store event data persistently 
* Purge older data


#### Notification Manager ####

This microservice is responsible for sending alerts and notifications:

* Send messages via SMS / push-notifications 

#### Reports Manager ####

This microservice is responsible for:

* Statistical reports
* Administrative user reports 
* Active/Inactive users

### Inter-service Communication ###

All services communicate between each other using gRPC. All services that are user-facing,
i.e. services that are sending data to users mobile phone are sending vis RESTful Json API.

