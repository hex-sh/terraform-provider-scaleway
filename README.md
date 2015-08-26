# terraform-provider-scaleway

Allows you to declaritvely describe your scaleway infrastructure within terraform.

We are planning to support:

* Servers
* Security groups
  - allows you to declare ASICs used for superfast hardware firewalls
* Volumes
  - Add multiple volumes to one server
* Snapshots
  - Create snapshots from volumes
* Images
  - Create bootable images from snapshots, which then can be used by servers
* S3-like object storage
  - Store as much data as you want for 0.02 cts per GB
