# hw-to-terraform

This is a personal project to help me learn more Go and hopefully transition to a SRE or Devops position. 

# Project Goals
- **Learn Go** through practical applications
- Collect server information from Unix and Windows systems (hostname, architecture, RAM, CPU, etc)
- Automatically generate **Infrastructure as Code (IaC)** files from this inventory data.
    - Support multiple IaC tools (Terraform, Bicep; more in future).
- Streamline the process of cloud infrastructure onboarding and environment replication.

# Features
 - **Cross-Platform** (Windows/Unix) inventory collection.
 - Converts system inventory data to:
    - **Terraform**
    - **Bicep**
- Workflow: Scan -> JSON -> IaC File

# Roadmap
### Implemented
* Collect system information and export to JSON

### To Be Implemented
* Generate Terraform file from inventory data
* Generate Bicep file from inventory data
* Support more resource types (networks, disks, tags)
