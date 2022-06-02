# CProtect

Copy protection and product activation utility written in Go.

## Introduction

Although software copy protection is an error-prone task to implement, this utility tries its best to accomplish that. Keywords:

* Activator - is a server or a person with license rights and activation capabilities.
* Request Code - is the computer's finger-print. A user attempting to install the product would get a request code. The user then would send this code to the activator.
* Activation Code - is the final code that will activate the application. The user will need to get this code from the activator server or person.
* Product - is the name of the software.
* Password - is a secure password held by the activator. Its length should be 16, 24, etc.

## API

To get the request code:

```go
reqCode, _ := cprotect.GetRequestCode("My Product")
```

To check if the product is installed in this PC:

```go
isInstalled, _ := cprotect.IsInstalled("My Product", <password>)
```

To check if an activation code is valid:

```go
valid, _ := cprotect.IsActivationCodeValid(<password>, <request-code>, <activation-code>)
```

To activate software using activation code:

```go
err := cprotect.Install("My Product", <activation-code>)
```