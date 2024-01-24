<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | EasyIO</div>

#  AB&T EasyIO

<dl>
  <dt>Type:</dt><dd>EasyIO</dd>
  <dt>Description:</dt><dd>Mixed I/O</dd>
  <dt>Vendor</dt><dd>AB&T</dd>
  <dt>Documentation</dt><dd><a href="http://www.bausano.net/">http://www.bausano.net/</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r0</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>Mixed I/O</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x0debacca</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0000002a</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=26 valign=top>TX PDOs</td>
<td><pre>0x1a00: Analog inputs</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:01  Input 1                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:02  Input 2                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:03  Input 3                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6000:04  Input 4                         INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: Digital input byte 1</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:01  Input 1                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:02  Input 2                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:03  Input 3                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:04  Input 4                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:05  Input 5                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:06  Input 6                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:07  Input 7                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6001:08  Input 8                         BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a02: Digital input byte 2</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:01  Input 1                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:02  Input 2                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:03  Input 3                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:04  Input 4                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:05  Input 5                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:06  Input 6                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:07  Input 7                         BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6002:08  Input 8                         BOOL</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a03: Digital output faults</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:01  Fault Byte 1                    BOOL</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6003:02  Fault Byte 2                    BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=21 valign=top>RX PDOs</td>
<td><pre>0x1600: Analog outputs</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7000:01  Output 1                        INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7000:02  Output 2                        INT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Digital output byte 1</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:01  Output 1                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:02  Output 2                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:03  Output 3                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:04  Output 4                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:05  Output 5                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:06  Output 6                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:07  Output 7                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7001:08  Output 8                        BOOL</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1602: Digital output byte 2</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:01  Output 1                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:02  Output 2                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:03  Output 3                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:04  Output 4                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:05  Output 5                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:06  Output 6                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:07  Output 7                        BOOL</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x7002:08  Output 8                        BOOL</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x0000079a" pid="0x0debacca" configPdos="true"&gt;
  &lt;syncManager idx="1" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="1" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="1" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="1" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
