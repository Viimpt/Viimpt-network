Viimpt 
====
Environment
---
> ubuntu 18.04  
> Docker 20.10.16  
> node.js 8.17.0  
> npm 5.6.0  
> go 1.11.1  

Hyperledger composer  
---
>Hyperledger Composer was a set of collaboration tools for building blockchain business networks that make it simple and fast for business owners and developers to create smart contracts and blockchain applications to solve business problems. Built with JavaScript, leveraging modern tools including node.js, npm, CLI and popular editors, Composer offered business-centric abstractions as well as sample apps with easy to test DevOps processes to create robust blockchain solutions that drive alignment across business requirements with technical development.[23]

>Blockchain package management tooling contributed by IBM. Composer was a user-facing rapid prototyping tooling, running on top of Hyperledger Fabric, which allows the easy management of Assets (data stored on the blockchain), Participants (identity management, or member services) and Transactions (Chaincode, a.k.a. Smart Contracts, which operate on Assets on the behalf of a Participant). The resulting application can be exported as a package (a BNA file) which may be executed on a Hyperledger Fabric instance, with the support of a Node.js application (based on the Loopback application framework) and provide a REST interface to external applications.

>Composer provided a GUI user interface "Playground" for the creation of applications, and therefore represents an excellent starting point for Proof of Concept work.


Hyperledger Indy 
---
>Hyperledger Indy provides tools, libraries, and reusable components for providing digital identities rooted on blockchains or other distributed ledgers so that they are interoperable across administrative domains, applications, and any other silo. Indy is interoperable with other blockchains or can be used standalone powering the decentralization of identity.  
>In Indy, proof of non-revocation is accomplished by having provers show that they can derive the value of the accumulator for their credential using a factor for the accumulator that they know, plus the product of all other factors. The verifier can see that the prover produces the right answer (because the answer is on the ledger), but does not know certain details of how the prover derived it. The issuer can revoke by changing the answer to the math problem in a way that defeats the prover.  


Hyperledger Fabric  
---
>Hyperledger Fabric is a permissioned blockchain infrastructure, originally contributed by IBM and Digital Asset, providing a modular architecture with a delineation of roles between the nodes in the infrastructure, execution of Smart Contracts (called "chaincode" in Fabric) and configurable consensus and membership services. A Fabric Network comprises (1) "Peer nodes", which execute chaincode, access ledger data, endorse transactions and interface with applications; (2) "Orderer nodes" which ensure the consistency of the blockchain and deliver the endorsed transactions to the peers of the network; and (3) Membership Service Providers (MSPs), each generally implemented as a Certificate Authority, managing X.509 certificates which are used to authenticate member identity and roles.[16] Hyperledger Fabric allows for use of different consensus algorithms, but the consensus algorithm that is most commonly used with the platform is Practical Byzantine Fault Tolerance (PBFT).  

>Fabric supports chaincode in Go and JavaScript (via Hyperledger Composer, or natively since v1.1) out-of-the-box, and other languages such as Java by installing appropriate modules. It is therefore potentially more flexible than competitors that only support a closed Smart Contract language.

