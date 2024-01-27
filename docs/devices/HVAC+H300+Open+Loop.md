<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | HVAC H300 Open Loop</div>

#  Control HVAC H300 Open Loop

<dl>
  <dt>Type:</dt><dd>HVAC H300 Open Loop</dd>
  <dt>Description:</dt><dd>HVAC H300 Open Loop</dd>
  <dt>Vendor</dt><dd>Control Techniques</dd>
  <dt>Documentation</dt><dd><a href=""></a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r512</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>HVAC H300 Open Loop</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x01010a02</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0200000c</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="Pump+Drive+F600+Open+Loop">Pump Drive F600 Open Loop r512</a><br/><a href="Unidrive+ID300+Open+Loop">Unidrive ID300 Open Loop r512</a><br/><a href="Unidrive+ID302+Open+Loop">Unidrive ID302 Open Loop r512</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=8 valign=top>TX PDOs</td>
<td><pre>0x1a00: Transmit PDO Mapping 1</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: Transmit PDO Mapping 2</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Modes of operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a05: Transmit PDO Mapping 6</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6044:00  vl velocity actual value        INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=8 valign=top>RX PDOs</td>
<td><pre>0x1600: Receive PDO Mapping 1</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Receive PDO Mapping 2</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1605: Receive PDO Mapping 6</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6042:00  vl target velocity              INT (16 bits)</pre></td>
</tr>
</table>
