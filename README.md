<!-- Improved compatibility of back to top link: See: https://github.com/othneildrew/Best-README-Template/pull/73 -->
<a name="readme-top"></a>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

<!-- PROJECT TITLE -->

<div align="center">
<h3 align="center">Deall ─ Backend Engineer Test</h3>
  <p align="center">
    Rest API CRUD User and User Login using Golang & MySQL
    <br />
    <br />
    <a href="https://github.com/aulkiller/Deall_BE_Test/issues">Report Bug</a>
    ·
    <a href="https://github.com/aulkiller/Deall_BE_Test/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
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
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

A Rest API Applicaton which allows CRUD User and User Login using Golang & MySQL. This API utilizes JWT to Authenticate it's endpoints and expires within 1 hour. There are two roles in this application, Admin and User. The Admin has access to all API CRUD, while the User only gets access to the user's data (Read). This application also implements Architecture Microservices using Kubernetes with Docker container deploy. 

Document related to this assignment can be seen [here](https://docs.google.com/document/d/1AfsaaEmpjgCMm3izfT0o3WUGHndWkZLl/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* Golang
* MySQL
* Docker
* Kubernetes

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

To get a local copy of this appplicaton running follow these steps.

### Prerequisites

Installed all of these container-related tools:
* Docker 
* Minikube
* Kubectl
  ```sh
  brew install kubectl
  ```
  or
    ```sh
  winget install -e --id Kubernetes.kubectl
  ```


### Installation

1. Get a free Dockerhub Account at [here](https://hub.docker.com)
2. Clone this repo
   ```sh
   git clone https://github.com/aulkiller/Deall_BE_Test.git
   ```
3. Install the Prequisites and login to docker via terminal
   ```sh
   docker login
   ```
4. Start the Minikube
   ```sh
   minikube start
   ```
5. Prepare the secret key and mysql pod deployment
   ```sh
   kubectl create -f mysql-secret.yaml
   kubectl apply -f mysql-db-pv.yaml
   kubectl apply -f mysql-db-pvc.yaml
   kubectl apply -f mysql-db-deployment.yaml
   kubectl apply -f mysql-db-service.yaml
   ```
6. Ensure the first pod successfully ran
   ```sh
   kubectl get pods
   ```
7. Build an image for the API applicaton. Replace YourApp with anything that you desired
   ```sh
   docker build -t <YourApp> .
   ```

   What i use:
      ```sh
   docker build -t app-kubernetes .
   ```
8. Tag and push the image into your dockerhub account
   ```sh
   docker tag <YourApp> <YourDockerHubName>/<RepoName>:<versions>
   docker push <YourDockerHubName>/<RepoName>:<versions>
   ```

   What i use:
   ```sh
   docker tag app-kubernetes aulkiller/deallapp:1.0.0
   docker push aulkiller/deallapp:1.0.0
   ```
9. Prepare the API pod deployment
   ```sh
   kubectl apply -f app-mysql-deployment.yaml
   kubectl apply -f app-mysql-service.yaml
   ```
10. Ensure the second pod and both services successfully ran
   ```sh
   kubectl get pods
   kubectl get services
   ```
11. Expose the API IP & Port from services/pods list
   ```sh
   minikube service <YourAPIService> --url
   ``` 
   or 
   ```sh
   kubectl port-forward <YourApiPods> 8080:8080
   ``` 

   What i use:
   ```sh
   kubectl port-forward deallapp-mysql-6fc6dc7456-jqwms 8080:8080
   ```  

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Use the exposed IP & Ports from step 9. If you use minikube service the API can be accessed on `http://Exposed-IP:Exposed-Port/` and If you use port-forwarding it should be on `http://127.0.0.1:8080/`.

Login credentials for admin (ID: 1, Role: Admin)
```yaml
{
    "username": "admin",
    "password": "passadmin"
}
```

Login credentials for auliaihza (ID: 2, Role: User)
```yaml
{
    "username": "auliaihza",
    "password": "pass"
}
```
API documentation by using Postman [here](https://docs.google.com/document/d/1AfsaaEmpjgCMm3izfT0o3WUGHndWkZLl/edit#)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Aulia Ihza Hendradi - auliaihza@gmail.com

Project Link: [https://github.com/aulkiller/Deall_BE_Test](https://github.com/aulkiller/Deall_BE_Test)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

