# SSL Simulation Protocol

---
**Note**

This protocol is still in draft mode and subject to changes

---

Protobuf files for a common simulation protocol for the Small Size League for easier exchange of different
simulator implementations:

* [ssl_simulation](./ssl_simulation.proto) - Main entry point of the protocol
* [ssl_simulation_config](./ssl_simulation_config.proto) - Change the configuration of the simulator
* [ssl_simulation_control](./ssl_simulation_control.proto) - Control the simulation itself
* [ssl_simulation_robot_control](./ssl_simulation_robot_control.proto) - Control the robots
* [ssl_simulation_robot_feedback](./ssl_simulation_robot_feedback.proto) - Receive robot feedback
* [ssl_gc_common](./ssl_gc_common.proto) - Common types from game-controller (RobotId) 
* [ssl_geometry](./ssl_geometry.proto) - Geometry from ssl-vision

The protocol defines mostly optional fields. All unset variables should be interpreted as unmodified.
A simulator may decide to not implement all features. Returning an error in case a field is set that is not
supported is considered good practice.
Errors should have a unique code for the simulator implementation that allows teams to automatically handle the
errors in their software.

## Communication

A simulator should offer a UDP port (default port to be defined) that a client connects to.
The client sends `SimulatorCommand` messages to the simulator, the simulator may answer with `SimulatorResponse` messages.

tbd: Request and response frequencies and response timing/synchronization

## Tournament Mode

When a simulator is used in a tournament, it should only allow robot control by teams.
Simulation control and configuration should be disabled in the simulator.
Optionally, it may only allow controlling one team color per IP.
