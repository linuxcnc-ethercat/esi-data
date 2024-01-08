<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | EpoCAT FR4000</div>

#  AB&T EpoCAT FR4000

<dl>
  <dt>Type:</dt><dd>EpoCAT FR4000</dd>
  <dt>Description:</dt><dd>FR4000 5 frequency controlled axes</dd>
  <dt>Vendor</dt><dd>AB&T Tecnologie Informatiche</dd>
  <dt>Documentation</dt><dd><a href="http://www.bausano.net/">http://www.bausano.net/</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r0</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>FR4000 5 frequency controlled axes</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x04decade</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00005a00</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=46 valign=top>TX PDOs</td>
<td><pre>0x1a00: Analog input</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:01  Input 1                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:02  Input 2                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: Pulse feedback</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:01  Drive 1 count                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:02  Drive 2 count                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:03  Drive 3 count                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:04  Drive 4 count                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:05  Drive 5 count                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a02: Pulse latch</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:01  Drive 1 latch                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:02  Drive 2 latch                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:03  Drive 3 latch                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:04  Drive 4 latch                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:05  Drive 5 latch                   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a06: Flag latch idx/probe</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6006:01  Drive 1 flag                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6006:02  Drive 2 flag                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6006:03  Drive 3 flag                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6006:04  Drive 4 flag                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6006:05  Drive 5 flag                    BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a04: Home microswitch</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6004:01  Drive 1 home                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6004:02  Drive 2 home                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6004:03  Drive 3 home                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6004:04  Drive 4 home                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6004:05  Drive 5 home                    BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a03: Drive status</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:01  Drive 1 ok                      BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:02  Drive 2 ok                      BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:03  Drive 3 ok                      BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:04  Drive 4 ok                      BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:05  Drive 5 ok                      BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:06  Auxiliary ok                    BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a05: Index</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6005:01  Drive 1 idx                     BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6005:02  Drive 2 idx                     BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6005:03  Drive 3 idx                     BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6005:04  Drive 4 idx                     BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6005:05  Drive 5 idx                     BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a07: Digital Input</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6007:01  Input 1                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6007:02  Input 2                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6007:03  Input 3                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6007:04  Input 4                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6007:05  Probe                           BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=43 valign=top>RX PDOs</td>
<td><pre>0x1600: Analog output</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7000:01  Output                          INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Frequency</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:01  Drive 1 freq                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:02  Drive 2 freq                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:03  Drive 3 freq                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:04  Drive 4 freq                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:05  Drive 5 freq                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: Enable</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:01  Drive 1 en                      BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:02  Drive 2 en                      BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:03  Drive 3 en                      BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:04  Drive 4 en                      BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:05  Drive 5 en                      BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1603: Direction</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7003:01  Drive 1 dir                     BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7003:02  Drive 2 dir                     BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7003:03  Drive 3 dir                     BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7003:04  Drive 4 dir                     BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7003:05  Drive 5 dir                     BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1604: Drive commands</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:01  Drive 1 res flag                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:02  Drive 2 res flag                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:03  Drive 3 res flag                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:04  Drive 4 res flag                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:05  Drive 5 res flag                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:06  Sel idx/probe                   BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7004:07  3state single ended             BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1605: Polarity</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7005:01  Drive 1 polarity                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7005:02  Drive 2 polarity                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7005:03  Drive 3 polarity                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7005:04  Drive 4 polarity                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7005:05  Drive 5 polarity                BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1606: Digital output</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:01  Drive 1 en 24V                  BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:02  Drive 2 en 24V                  BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:03  Drive 3 en 24V                  BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:04  Drive 4 en 24V                  BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:05  Drive 5 en 24V                  BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:06  Auxiliary en 24V                BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:07  Output 1 24V                    BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7006:08  Output 2 24V                    BOOL</pre></td>
</tr>
</table>
