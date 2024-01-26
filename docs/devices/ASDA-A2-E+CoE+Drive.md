<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | ASDA-A2-E CoE Drive</div>

#  Delta ASDA-A2-E CoE Drive

<dl>
  <dt>Type:</dt><dd>ASDA-A2-E CoE Drive</dd>
  <dt>Description:</dt><dd>Delta ASDA-A2-E EtherCAT(CoE) Drive</dd>
  <dt>Vendor</dt><dd>Delta Electronics, Inc.</dd>
  <dt>Documentation</dt><dd><a href=""></a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r516</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>Delta ASDA-A2-E EtherCAT(CoE) Drive</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x10305070</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x02040608</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=17 valign=top>TX PDOs</td>
<td><pre>0x1a00: 1st TxPDO Mapping</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  ActualPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x606c:00  Velocity actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6077:00  ActualTorque                    INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  ModeOfOperationDisplay          SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: 2nd TxPDO Mapping</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  ActualPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a02: 3rd TxPDO Mapping</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  ActualPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x606c:00  Velocity actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a03: 4th TxPDO Mapping</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  ActualPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6077:00  ActualTorque                    INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=15 valign=top>RX PDOs</td>
<td><pre>0x1600: 1st RxPDO Mapping</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  TargetPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60ff:00  TargetVelocity                  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6071:00  TargetTorque                    INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  ModeOfOperation                 SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: 2nd RxPDO Mapping</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  TargetPosition                  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: 3rd RxPDO Mapping</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60ff:00  TargetVelocity                  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1603: 4th RxPDO Mapping</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6071:00  TargetTorque                    INT (16 bits)</pre></td>
</tr>
</table>
