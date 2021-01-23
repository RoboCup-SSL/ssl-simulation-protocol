# SSL Simulation Protocol

---
**Note**

This protocol is still in draft mode and subject to changes

---

Protobuf files for a common simulation protocol for the Small Size League for easier exchange of different
simulator implementations.

It consists of the following files:

* [ssl_simulation_robot_control](./ssl_simulation_robot_control.proto) - Control the robots
* [ssl_simulation_robot_feedback](./ssl_simulation_robot_feedback.proto) - Receive robot feedback
* [ssl_simulation_control](./ssl_simulation_control.proto) - Control the simulation itself
* [ssl_simulation_config](./ssl_simulation_config.proto) - Change the configuration of the simulator
* [ssl_simulation_error](./ssl_simulation_error.proto) - Error messages for responses
* [ssl_gc_common](./ssl_gc_common.proto) - Common types from game-controller (RobotId) 
* [ssl_geometry](./ssl_geometry.proto) - Geometry from ssl-vision

The protocol defines mostly optional fields. All unset variables should be interpreted as unmodified.
A simulator may decide to not implement all features. Returning an error in case a field is set that is not
supported is considered good practice.
Errors should have a unique code for the simulator implementation that allows teams to automatically handle the
errors in their software.

## Communication

There are three ports that a simulator should offer:

1. Simulation Control (10300): Accepts `SimulatorCommand` messages and returns `SimulatorResponse` messages.
1. Robot Control Blue (10301): Accepts `RobotControl` messages and returns `RobotControlResponse` messages, only for blue team
1. Robot Control Yellow (10302): Accepts `RobotControl` messages and returns `RobotControlResponse` messages, only for yellow team

All connections use bidirectional UDP communication.

## Tournament Mode

When a simulator is used in a tournament, it should only allow one connection per port.
Simulation control and configuration should be disabled in the simulator or for example restricted to localhost
so that teams can not (accidentally) connect to the simulation control.

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
