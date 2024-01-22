<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | SGD7S-xxxDA0xxxxF64 CoE Drive</div>

#  Yaskawa SGD7S-xxxDA0xxxxF64 CoE Drive

<dl>
  <dt>Type:</dt><dd>SGD7S-xxxDA0xxxxF64 CoE Drive</dd>
  <dt>Description:</dt><dd>SGD7S-xxxDA0xxxxF64 EtherCAT(CoE) SERVOPACK Rev8.04</dd>
  <dt>Vendor</dt><dd>Yaskawa Electric Corporation</dd>
  <dt>Documentation</dt><dd><a href=""></a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 4 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td  colspan=4 align="center"><pre>r8</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>SGD7S-xxxDA0xxxxF64 EtherCAT(CoE) SERVOPACK Rev8.04</pre></td>
<td ><pre>SGD7S-xxxDA0xxxxF64 EtherCAT(CoE) SERVOPACK Rev8.10</pre></td>
<td ><pre>SGD7S-xxxDA0xxxxF64 EtherCAT(CoE) SERVOPACK Rev8.14</pre></td>
<td ><pre>SGD7S-xxxDA0xxxxF64 EtherCAT(CoE) SERVOPACK Rev8.19</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=4 align="center"><pre>0x02200401</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00080004</pre></td>
<td ><pre>0x00080010</pre></td>
<td ><pre>0x00080014</pre></td>
<td ><pre>0x00080019</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td  colspan=4 align="center"></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=18 valign=top>TX PDOs</td>
<td colspan=4 align="left"><pre>0x1a00: 1st Transmit PDO mapping</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6064:00  Position actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6077:00  Torque actual value             INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x60f4:00  Following error actual value    DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6061:00  Modes of operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x60b9:00  Touch probe status              UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x60ba:00  Touch probe 1 position value    DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=4 align="left"><pre>0x1a01: 2nd Transmit PDO mapping</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6064:00  Position actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=4 align="left"><pre>0x1a02: 3rd Transmit PDO mapping</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6064:00  Position actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=4 align="left"><pre>0x1a03: 4th Transmit PDO mapping</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6064:00  Position actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=4 align="left"><pre>  0x6077:00  Torque actual value             INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=17 valign=top>RX PDOs</td>
<td colspan=4 align="left"><pre>0x1600: 1st Receive PDO mapping</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x607a:00  Target position                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x60ff:00  Target velocity                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6071:00  Target torque                   INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6072:00  Max torque                      UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6060:00  Modes of operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x60b8:00  Touch probe function            UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=4 align="left"><pre>0x1601: 2nd Receive PDO mapping</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x607a:00  Target position                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=4 align="left"><pre>0x1602: 3rd Receive PDO mapping</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x60ff:00  Target velocity                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=4 align="left"><pre>0x1603: 4th Receive PDO mapping</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=4 align="left"><pre>  0x6071:00  Target torque                   INT (16 bits)</pre></td>
</tr>
</table>
