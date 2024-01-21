<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | FB1311 Dig. In/Out</div>

#  Beckhoff FB1311 Dig. In/Out

<dl>
  <dt>Type:</dt><dd>FB1311 Dig. In/Out</dd>
  <dt>Description:</dt><dd>FB1311 16 Ch. Dig. In-/Output 2xMII</dd>
  <dt>Vendor</dt><dd>Beckhoff Automation GmbH & Co. KG</dd>
  <dt>Documentation</dt><dd><a href="http://www.beckhoff.com/FB1311">http://www.beckhoff.com/FB1311</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r100</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>FB1311 16 Ch. Dig. In-/Output 2xMII</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x051f0862</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0064008e</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="FB1111+Dig.+In%2FOut">FB1111 Dig. In/Out r100,r101</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=2 valign=top>TX PDOs</td>
<td><pre>0x1600: Byte 0</pre></td>
<td></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1601: Byte 1</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=2 valign=top>RX PDOs</td>
<td><pre>0x1a00: Byte 0</pre></td>
<td></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1a01: Byte 1</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000002" pid="0x051f0862" configPdos="true"&gt;
  &lt;syncManager idx="2" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="2" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
