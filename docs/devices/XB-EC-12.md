<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | XB-EC-12</div>

#  Bosch XB-EC-12

<dl>
  <dt>Type:</dt><dd>XB-EC-12</dd>
  <dt>Description:</dt><dd>XB-EC-12 Bus coupler PwrIn UL UP</dd>
  <dt>Vendor</dt><dd>Bosch Rexroth AG</dd>
  <dt>Documentation</dt><dd><a href="https://dc-qr.com?m=R911406090">https://dc-qr.com?m=R911406090</a></dd>
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
<td  colspan=2 align="center"><pre>XB-EC-12 Bus coupler PwrIn UL UP</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=2 align="center"><pre>0x00242a0f</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x00000100</pre></td>
<td ><pre>0x00010100</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="XC811201">XC811201 r0</a></pre></td>
<td ></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=13 valign=top>TX PDOs</td>
<td colspan=2 align="left"><pre>0x1a00: UP Supply periphery</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6000:01  UP Voltage                      UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6000:02  UP Current                      UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a01: UL Supply logic</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6010:01  UL Voltage                      UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6010:02  UL Current                      UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td  colspan=2 align="left"><pre>0x1a02: State</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:01  UP Undervoltage                 BOOL</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:02  UP Overvoltage                  BOOL</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:03  UP Overcurrent                  BOOL</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:04  UL Undervoltage                 BOOL</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:05  UL Overvoltage                  BOOL</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6020:06  UL Overcurrent                  BOOL</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000024" pid="0x00242a0f" configPdos="true"&gt;
  &lt;syncManager idx="3" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="3" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="3" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
