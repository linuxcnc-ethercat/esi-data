<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | ECT60X2</div>

#  Shenzhen ECT60X2

<dl>
  <dt>Type:</dt><dd>ECT60X2</dd>
  <dt>Description:</dt><dd>ECT60X2(COE)</dd>
  <dt>Vendor</dt><dd>Shenzhen Ruitech Mechanical and Electrical Technology Co.,Ltd. </dd>
  <dt>Documentation</dt><dd><a href="http://www.szruitech.com/">http://www.szruitech.com/</a></dd>
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
<td ><pre>ECT60X2(COE)</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x0a880006</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00000202</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="ECR60X2">ECR60X2 r0</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=20 valign=top>TX PDOs</td>
<td><pre>0x1a00: CHN 1 Transmit PDO 1</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Modes of Operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  Position Actual Value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60fd:00  Digital Inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: CHN 1 Transmit PDO 2</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Modes of Operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x606c:00  Velocity Actual Value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60fd:00  Digital Inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a10: CHN 2 Transmit PDO 1</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6841:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6861:00  Modes of Operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6864:00  Position Actual Value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x68fd:00  Digital Inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a11: CHN 2 Transmit PDO 2</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6841:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6861:00  Modes of Operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x686c:00  Velocity Actual Value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x68fd:00  Digital Inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=34 valign=top>RX PDOs</td>
<td><pre>0x1600: CHN 1 Receive PDO 1</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  ModeOfOperation                 SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  Profile Target Position         DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: CHN 1 Receive PDO 2</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  Profile Target Position         DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6081:00  Profile Velocity                UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6083:00  Profile Target Acceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6084:00  Profile Target Deceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of Operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: CHN 1 Receive PDO 3</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6083:00  Profile Target Acceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6084:00  Profile Target Deceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60ff:00  Target Velocity                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of Operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1610: CHN 2 Receive PDO 1</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6840:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6860:00  ModeOfOperation                 SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x687a:00  Profile Target Position         DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1611: CHN 2 Receive PDO 2</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6840:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x687a:00  Profile Target Position         DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6881:00  Profile Velocity                UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6883:00  Profile Target Acceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6884:00  Profile Target Deceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6860:00  Modes of Operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1612: CHN 2 Receive PDO 3</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6840:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6883:00  Profile Target Acceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6884:00  Profile Target Deceleration     UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x68ff:00  Target Velocity                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6860:00  Modes of Operation              SINT (8 bits)</pre></td>
</tr>
</table>
