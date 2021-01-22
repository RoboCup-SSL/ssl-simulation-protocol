# SSL Simulation Protocol

---
**Note**

This protocol is still in draft mode and subject to changes

---

Protobuf files for a common simulation protocol for the Small Size League for easier exchange of different
simulator implementations.

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
