<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">gRPC + Golang microservice PoC</h3>

  <p align="center">
    A simple idea for microservice implementation using Golang and gRPC(+protobuf)
    <br />
    <a href="https://github.com/emilianozublena/microservices"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/emilianozublena/microservices/issues">Report Bug</a>
    ·
    <a href="https://github.com/emilianozublena/microservices/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

This project was created with the sole purpose of learning. I come from typical object-oriented synchronous language, so moving away from that into the Golang world, came with a lot of challenges.
This project was a mere excuse for me to understand topics surrounding golang (such as Types, Concurrency, etc) while trying to understand how does Golang suggest to implement things that i'm used to work with in an object-oriented language (such as polymorphism, just to name one).

There are A LOT of these examples, and this is just a humble attempt of my own, by no means this should be used in production, nor should be understood as *the best way to achieve a ms arch...*

### Built With

* [Golang](https://golang.org/)
* [MongoDB](https://mongodb.com/)
* [Bongo ODM](https://github.com/go-bongo/bongo)
* [Routific API](https://routific.com/)


<!-- GETTING STARTED -->
## Getting Started

You can either choose to build the go package or just simply run it by using the *go run* command. Choice is yours.
Keep in mind that under /internal/client you have an example gRPC client built in Golang just for testing the service. The client is not tested and is not intended to be, as tipically the microservice wouldn't be holding a client inside of itself.

### Prerequisites

Follow installation for [Golang](https://golang.org/), [gRPC and protobuf](https://grpc.io/)

<!-- USAGE EXAMPLES -->
## Usage

Start the server directly:
```go run main.go
```

Or either build it and run it
```go build
./microservices
```

After server is running, you can play with your own client, or just use the one shipped here.
The client will accept os.Args, in which you can tell the client what actions should it be taking
```cd internal/client
go run main.go create
```

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Emiliano Zublena - [@emilianozublena](https://www.linkedin.com/in/emilianozublena/) - ezublena@gmail.com

Project Link: [https://github.com/emilianozublena/microservices](https://github.com/emilianozublena/microservices)


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/emilianozublena/microservices.svg?style=for-the-badge
[contributors-url]: https://github.com/emilianozublena/microservices/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/emilianozublena/microservices.svg?style=for-the-badge
[forks-url]: https://github.com/emilianozublena/microservices/network/members
[stars-shield]: https://img.shields.io/github/stars/emilianozublena/microservices.svg?style=for-the-badge
[stars-url]: https://github.com/emilianozublena/microservices/stargazers
[issues-shield]: https://img.shields.io/github/issues/emilianozublena/microservices.svg?style=for-the-badge
[issues-url]: https://github.com/emilianozublena/microservices/issues
[license-shield]: https://img.shields.io/github/license/emilianozublena/microservices.svg?style=for-the-badge
[license-url]: https://github.com/emilianozublena/microservices/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/othneildrew
[product-screenshot]: images/screenshot.png