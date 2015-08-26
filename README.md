programming exercise:

Theoretical situation:
We have data coming from a security device that outputs json. We need to monitor a folder for new files and process them as as soon as possible.
The output should be a status of the system with statistics every
second. The counts should be totals for that time interval. The avg should be for the time interval not the total for all processing. If nothing is inserted into the system you should output all 0's.

The “processing time” for a single input file is defined as the time between the file appearing, and the time it takes for that file to be read/parsed and included in the counts.  The average time reported should include all the messages counted in that output line.  It’s okay to state your assumptions, as parts of the exercise are intentionally open-ended.

The scope of this exercise is intended to be contained to a few hours of development at most, but however small please treat the solution as you would treat clean, production code that you would produce as part of a team at the end of a release to a high-performance production environment. We’re looking for a demonstration of development practices that you would employ in such a real-world setting.

Example:

—-existing file
{"Type":Door, "Date":"2014-02-01 10:01:02", "open": true}

--new file
{"Type":Alarm, "Date":"2014-02-01 10:01:01", "name":"fire", "floor":"1", "Room": "101"}

Output:
“DoorCnt: 0, ImgCnt:0, AlarmCnt:1, avgProcessingTime: 10ms"


Some example input files should be attached.
