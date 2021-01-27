# SSL Simulation Protocol

---
**Note**

This protocol is still in draft mode and subject to changes

---

Protobuf files for a common simulation protocol for the Small Size League for easier exchange of different
simulator implementations.

It consists of the following files:

* [ssl_simulation_robot_control](./proto/ssl_simulation_robot_control.proto) - Control the robots
* [ssl_simulation_robot_feedback](./proto/ssl_simulation_robot_feedback.proto) - Receive robot feedback
* [ssl_simulation_control](./proto/ssl_simulation_control.proto) - Control the simulation itself
* [ssl_simulation_config](./proto/ssl_simulation_config.proto) - Change the configuration of the simulator
* [ssl_simulation_error](./proto/ssl_simulation_error.proto) - Error messages for responses
* [ssl_gc_common](./proto/ssl_gc_common.proto) - Common types from game-controller (RobotId) 
* [ssl_geometry](./proto/ssl_vision_geometry.proto) - Geometry from ssl-vision

The protocol defines mostly optional fields. All unset variables should be interpreted as unmodified.
A simulator may decide to not implement all features. Returning an error in case a field is set that is not
supported is considered good practice.
Errors should have a unique code for the simulator implementation that allows teams to automatically handle the
errors in their software.

## Communication

There are three ports that a simulator should offer:

### Simulation Control
Control and configure the simulation.

* Accepts `SimulatorCommand` messages ([ssl_simulation_control](./proto/ssl_simulation_control.proto))
* Returns `SimulatorResponse` messages ([ssl_simulation_control](./proto/ssl_simulation_control.proto))
* Default port: 10300 UDP

### Robot Control Blue
Control the blue team.

* Accepts `RobotControl` messages ([ssl_simulation_robot_control](./proto/ssl_simulation_robot_control.proto))
* Returns `RobotControlResponse` messages ([ssl_simulation_robot_feedback](./proto/ssl_simulation_robot_feedback.proto))
* Default port: 10301 UDP

### Robot Control Yellow
Control the yellow team.

* Accepts `RobotControl` messages ([ssl_simulation_robot_control](./proto/ssl_simulation_robot_control.proto))
* Returns `RobotControlResponse` messages ([ssl_simulation_robot_feedback](./proto/ssl_simulation_robot_feedback.proto))
* Default port: 10302 UDP

All connections use bidirectional UDP communication.

## Tournament Mode

When a simulator is used in a tournament, it should only allow UDP packets from one IP per team.
Simulation control and configuration should be disabled in the simulator or for example restricted to localhost
or also to only one IP, so that teams can not (accidentally) connect to the simulation control, but
an autoRef or similar can.

## Synchronous communication

During development and for automated tests it might be useful to have a synchronous communication with the team software.
Additionally, a multicast protocol might cause issues in these scenarios.
There are also things like authentication for a tournament mode to be considered.

For that, an additional communication interface is planned. There are multiple ideas:
1. Use a bidirectional TCP connection, like the game-controller provides (messages a drafted in [ssl_simulation_synchronous.proto](./proto/ssl_simulation_synchronous.proto))
1. Providing C-Interfaces to avoid network communication completely
1. Using gRPC, esp. for configuration

In any case, the protobuf messages would still be used, but a layer would be added on top.
These different ideas still need more evaluation. Feedback and opinions can be addressed to the SSL TC.

## Design decisions

### Multiple ports
The protocol has been divided into three individual connections to easily disable the control protocol during
a tournament and to make sure each team can only command their own robots and can only receive feedback from their
own robots.

### Custom vs. common messages
It is expected that each simulator has their specifics, and it will be hard to generalize the protocol in a way,
that all parameters can be handled by all implementations and even that they have the exact same behavior.
For that reason, each simulator can add custom messages for realism, robot specs and robot feedback with
an individual documentation and behavior.

### Configuration in simulator vs. in protocol
All simulators will require some configuration. While there are some with a UI where the parameters can be
configured, others may be completely headless and stateless.
Additionally, teams may want to integrate a simulator in their own software and configure it from there, instead
of running a separate UI.
That's why the simulator configuration can be specified via the protocol.

### Use UDP for all connections
UDP is an unreliable communication protocol, so datagrams (and thus protobuf messages) may get lost.
This is ok for robot commands, as they are sent repeatedly in a high frequency anyway and there is no need for 
a retransmission.
It is also a little simpler to send data through UDP than through TCP (implementation wise).
To keep the overall protocol simple, the simulation control also uses UDP.
