<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | FB1111 Dig. In</div>

#  Beckhoff FB1111 Dig. In

<dl>
  <dt>Type:</dt><dd>FB1111 Dig. In</dd>
  <dt>Description:</dt><dd>FB1111 32 Ch. Dig. Input 2xMII</dd>
  <dt>Vendor</dt><dd>Beckhoff Automation GmbH & Co. KG</dd>
  <dt>Documentation</dt><dd><a href="http://www.beckhoff.com/FB1111">http://www.beckhoff.com/FB1111</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 2 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r400</pre></td>
<td ><pre>r401</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td  colspan=2 align="center"><pre>FB1111 32 Ch. Dig. Input 2xMII</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=2 align="center"><pre>0x04570862</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0190008e</pre></td>
<td ><pre>0x0191008e</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td  colspan=2 align="center"><pre><a href="EL9800">EL9800 r1</a><br/><a href="FB1311+Dig.+In">FB1311 Dig. In r400</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=4 valign=top>TX PDOs</td>
<td colspan=2 align="left"><pre>0x1600: Byte 0</pre></td>
<td></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1601: Byte 1</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1602: Byte 2</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1603: Byte 3</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000002" pid="0x04570862" configPdos="true"&gt;
  &lt;syncManager idx="0" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="0" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="0" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="0" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
