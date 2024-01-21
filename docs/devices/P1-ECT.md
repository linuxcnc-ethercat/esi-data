<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | P1-ECT</div>

#  Hitachi P1-ECT

<dl>
  <dt>Type:</dt><dd>P1-ECT</dd>
  <dt>Description:</dt><dd>P1-ECT EtherCAT Communication Unit for SJ-P1</dd>
  <dt>Vendor</dt><dd>Hitachi Industrial Equipment Systems Co.,Ltd.</dd>
  <dt>Documentation</dt><dd><a href="http://www.hitachi-ies.co.jp/english/">http://www.hitachi-ies.co.jp/english/</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 2 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r0</pre></td>
<td ><pre>r1</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td  colspan=2 align="center"><pre>P1-ECT EtherCAT Communication Unit for SJ-P1</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=2 align="center"><pre>0x03010100</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00000100</pre></td>
<td ><pre>0x00010001</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td  colspan=2 align="center"></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=4 valign=top>TX PDOs</td>
<td colspan=2 align="left"><pre>0x1a05: 6th transmit PDO Mapping</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6041:00  Statusword                      UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6043:00  vl velocity demand              INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a00: 1st transmit PDO Mapping</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=4 valign=top>RX PDOs</td>
<td colspan=2 align="left"><pre>0x1605: 6th receive PDO Mapping</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td  colspan=2 align="left"><pre>  0x6040:00  Controlword                     UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=2 align="left"><pre>  0x6042:00  vl target velocity              INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td  colspan=2 align="left"><pre>0x1600: 1st receive PDO Mapping</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x0000051d" pid="0x03010100" configPdos="true"&gt;
  &lt;syncManager idx="3" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
