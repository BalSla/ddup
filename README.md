# ddup
Simple console-based application specifically designed to remove file duplicates from a directory based on the contents of another directory.

## Usage
```shell
ddup -source <source_dir> -target <target_dir>
```
To simulate use:
```shell
ddup -source <source_dir> -target <target_dir> -dryrun
```


## Specialization
There are some similar scripts in GitHib. However, the goal of this utility is to speed up cleaning of the directory containg source files which already backed up by DaVinci Resolve (DR) inside `source` dir.
The problem is if files are backed up by DR inside MediaFiles directory with structure of their parent directories

- project dir
  - file                       
  - file

- MediaFiles
  - file
  - file

but in following case

- A
  - B (project dir)
    - file a
    - file a
    - project dir
- C
  - D
    - file c

structure will be different

- MediaFiles
  - A
    - B
      - file a
      - file a
  - C
    - D
      - file c 
     
Thus utility tryes to find...
