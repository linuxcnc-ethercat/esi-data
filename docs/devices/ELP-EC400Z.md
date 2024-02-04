<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | ELP-EC400Z</div>

#  Leadshine ELP-EC400Z

<dl>
  <dt>Type:</dt><dd>ELP-EC400Z</dd>
  <dt>Description:</dt><dd>ELP-EC400Z(COE)</dd>
  <dt>Vendor</dt><dd>Leadshine Technology Co.,Ltd.</dd>
  <dt>Documentation</dt><dd><a href="http://www.leadshine.com">http://www.leadshine.com</a></dd>
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
<td ><pre>ELP-EC400Z(COE)</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x000000a2</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00000001</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="ELP-EC1000Z">ELP-EC1000Z r0</a><br/><a href="ELP-EC100Z">ELP-EC100Z r0</a><br/><a href="ELP-EC1500Z">ELP-EC1500Z r0</a><br/><a href="ELP-EC2000Z">ELP-EC2000Z r0</a><br/><a href="ELP-EC200Z">ELP-EC200Z r0</a><br/><a href="ELP-EC3000Z">ELP-EC3000Z r0</a><br/><a href="ELP-EC750Z">ELP-EC750Z r0</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=9 valign=top>TX PDOs</td>
<td><pre>0x1a00: Transmit PDO 1</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x603f:00  Last error code                 UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Modes of operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  Actual motor position           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60b9:00  Touch Probe Status              UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60ba:00  Touch probe pos1 pos value      DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60fd:00  Digital inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: Transmit PDO 2</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=22 valign=top>RX PDOs</td>
<td><pre>0x1600: Receive PDO 1</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  Profile target position         DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60b8:00  Touch Probe Function            UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Receive PDO 2</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60ff:00  Target velocity                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60b2:00  Torque offset                   INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              USINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: Receive PDO 3</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6071:00  Target Torque                   INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6087:00  Torque slope                    UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              USINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1603: Receive PDO 4</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6098:00  Homing method                   SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6099:01  Homing velocity (fast)          UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6099:02  Homing velocity (slow)          UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x609a:00  Homing acceleration             UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607c:00  Homing offset                   DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              USINT (8 bits)</pre></td>
</tr>
</table>