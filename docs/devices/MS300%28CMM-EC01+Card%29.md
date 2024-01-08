<div class="nav"><a href="/esi-data">ESI data</a> | <a href="/esi-data/devices">devices</a> | MS300(CMM-EC01 Card)</div>

#  Delta MS300(CMM-EC01 Card)

<dl>
  <dt>Type:</dt><dd>MS300(CMM-EC01 Card)</dd>
  <dt>Description:</dt><dd>Delta MS300 EtherCAT(CoE)</dd>
  <dt>Vendor</dt><dd>Delta Electronics, Inc.</dd>
  <dt>Documentation</dt><dd><a href=""></a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r1</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>Delta MS300 EtherCAT(CoE)</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x10400200</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00010002</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="MS300%28CMM-EC02+Card%29">MS300(CMM-EC02 Card) r2</a><br/><a href="MS300%28CMM-EC02+Card%29">MS300(CMM-EC02 Card) r2</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=7 valign=top>TX PDOs</td>
<td><pre>0x1a00: 1st TxPDO Mapping</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status Word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6043:00  vl velocity demand              INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Mode Of Operation Display       SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: 2nd TxPDO Mapping</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a02: 3rd TxPDO Mapping</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a03: 4th TxPDO Mapping</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=7 valign=top>RX PDOs</td>
<td><pre>0x1600: 1st RxPDO Mapping</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control Word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6042:00  vl target velocity              INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Mode Of Operation               SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: 2nd Receive PDO mapping</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: 3nd Receive PDO mapping</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1603: 4nd Receive PDO mapping</pre></td>
</tr>
</table>
