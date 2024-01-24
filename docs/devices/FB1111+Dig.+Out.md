<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | FB1111 Dig. Out</div>

#  Beckhoff FB1111 Dig. Out

<dl>
  <dt>Type:</dt><dd>FB1111 Dig. Out</dd>
  <dt>Description:</dt><dd>FB1111 32 Ch. Dig. Output 2xMII</dd>
  <dt>Vendor</dt><dd>Beckhoff Automation GmbH & Co. KG</dd>
  <dt>Documentation</dt><dd><a href="http://www.beckhoff.com/FB1111">http://www.beckhoff.com/FB1111</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 2 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r500</pre></td>
<td ><pre>r501</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td  colspan=2 align="center"><pre>FB1111 32 Ch. Dig. Output 2xMII</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=2 align="center"><pre>0x04570862</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x01f4008e</pre></td>
<td ><pre>0x01f5008e</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td  colspan=2 align="center"><pre><a href="FB1311+Dig.+Out">FB1311 Dig. Out r500</a></pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=4 valign=top>RX PDOs</td>
<td colspan=2 align="left"><pre>0x1a00: Byte 0</pre></td>
<td></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a01: Byte 1</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a02: Byte 2</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a03: Byte 3</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000002" pid="0x04570862" configPdos="true"&gt;
&lt;/slave&gt;
</pre>
