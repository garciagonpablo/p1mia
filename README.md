# EXT3 Recreation Project
- Install the DOT (Graphviz) command in your console.
- Install golang, and install the project.
- Execute main.go

## Example of the commands available:

### Parameters available:
Size parameters:
- M (Megabytes)
- K (Kilobytes)
- B (bytes)
Fit Parameters:
- FF (First fit)
- BF (Best fit)
- WF (Worst fit)
Partition type:
- P (Primary)
- L (Logic)
- E (Extended)

  
### Disk Creation:
Example:
``Mkdisk -size=60 -unit=M -fit=FF                   # 60M A``

### Partition Creation:
Maximum of 4 partitions, an extended partition can have 4 logic partitions.
```
fdisk -type=P -unit=b -name=Part1 -size=20971520 -driveletter=A -fit=BF    # 20M
fdisk -type=P -unit=k -name=Part2 -size=10240 -driveletter=A -fit=BF       # 10M
fdisk -type=P -unit=M -name=Part3 -size=10 -driveletter=A -fit=BF          # 10M
fdisk -type=E -unit=b -name=Part4 -size=10485760 -driveletter=A -fit=BF    # 10M
fdisk -type=L -unit=b -name=LogicPart -size=4485760 -driveletter=A -fit=BF # 4M

```
## Partition Mount:
The commented ID can be used to display reports. The ids of the partitions are given by: (Driveletter of the disk)(Number of partition mounted in the disk starting at 1)07.
```
mount -driveletter=A -name=Part1 #A107
mount -driveletter=A -name=Part2 #A207
```

### Reports:
Tree File report and Superblock report.
```
rep -id=A107 -path=~/EXT3/report1_sb.jpg -name=sb
rep -id=A107 -path=~/EXT3/report1_tree.jpg -name=tree
```

### File System Creation:
Creates the selected File System in the id specified.
```
mkfs -type=full -id=A107 -fs=2fs
mkfs -type=full -id=B107 -fs=3fs
```
### Login/Logout:
Default user to login in the specified mount.
```
login -user=root -pass=123 -id=A107
logout
```
