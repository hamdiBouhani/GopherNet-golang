# Architecture project

## Objective

> This assignment aims to evaluate your understanding and proficiency in cloudarchitecture. You should demonstrate your ability to design a robust, scalable, andefficient cloud solution for a hypothetical burrow advertising platform, GopherNet.


GopherNet is an emerging digital platform that facilitates advertising for gopherburrows, serving thousands of users daily. It has multiple components, including user registration and authentication, search functionality and burrow ad placement.GopherNet is looking to transition from it’s current traditional server setup to ascalable and efficient cloud-based infrastructure.Design an architecture to support the GopherNet platform. Your design should usecloud services from your preferred cloud provider (eg, AWS, GCP)Your solution should ensure high availability, scalability, and cost-effectiveness.Task 1Draw a detailed diagram of your proposed cloud architecture. The diagram shouldclearly illustrate how different components of the architecture interact with eachother. You may use any diagramming tool of your choice.Task 2Write a paragraph for the following topics:How to manage and provision production, test and dev environmentsBasic explanation of how deployments will workHigh-level explanation of the local development workflowTask 3In a few small paragraphs - discuss how your cloud architecture could handle asudden surge in traffic, which might occur due to a promotional event. Explain howyour design maintains high performance and availability during these peak times.DeliverableYour final submission should include your architecture diagram and writtenexplanations. Please package these into a single PDF document.We look forward to your creative solutions.


## Task 1

Draw a detailed diagram of your proposed cloud architecture. The diagram should clearly illustrate how different components of the architecture interact with eachother. You may use any diagramming tool of your choice.


### application Architecture
![proposed Architecture](./architecture-img.png)
### cloud Architecture
![proposed Architecture](./cloud_architecture.png)

## Task 2
> How to manage and provision production, test and dev environments.

for each Environment I will create a separate VPC, and a separate GKE cluster. I will use terraform to provision the infrastructure. will create a terraform module for each environment, and use terraform workspaces to manage the environments.

* Basic explanation of how deployments will workHigh-level explanation of the local development workflow

I will use github actions to build the docker images and push them to docker hub. I will use terraform to deploy the application to the different environments. I will use helm to deploy the application to the kubernetes cluster. by using helm terraform provider I will be able to deploy the application to the different environments. Or we can use github actions to deploy the application to the different environments.



## Task 3
In a few small paragraphs - discuss how your cloud architecture could handle asudden surge in traffic, which might occur due to a promotional event. Explain how your design maintains high performance and availability during these peak times.

to scale the application horizontally, I will use a load balancer to distribute the traffic between the different pods. I will use a horizontal pod autoscaler to scale the pods based on the cpu usage. I will use a cluster autoscaler to scale the cluster based on the number of pending pods. I will use a vertical pod autoscaler to scale the pods based on the cpu usage. I will use a node autoscaler to scale the nodes based on the number of pending pods. I will use a pod disruption budget to make sure that there is always a minimum number of pods running. I will use a pod anti affinity to make sure that the pods are not scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node. I will use a pod affinity to make sure that the pods are scheduled on the same node.