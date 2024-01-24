<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | EX260-SEC1</div>

#  SMC EX260-SEC1

<dl>
  <dt>Type:</dt><dd>EX260-SEC1</dd>
  <dt>Description:</dt><dd>EX260-SEC1 EtherCAT SI Unit (32DO_PNP) Rev1.2</dd>
  <dt>Vendor</dt><dd>SMC Corporation</dd>
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
<td ><pre>EX260-SEC1 EtherCAT SI Unit (32DO_PNP) Rev1.2</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x01000001</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00010002</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="EX260-SEC2">EX260-SEC2 r1</a></pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=4 valign=top>RX PDOs</td>
<td><pre>0x1600: Byte 0</pre></td>
<td></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Byte 1</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: Byte 2</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1603: Byte 3</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000114" pid="0x01000001" configPdos="true"&gt;
&lt;/slave&gt;
</pre>
