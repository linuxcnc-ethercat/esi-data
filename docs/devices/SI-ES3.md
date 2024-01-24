<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | SI-ES3</div>

#  Yaskawa SI-ES3

<dl>
  <dt>Type:</dt><dd>SI-ES3</dd>
  <dt>Description:</dt><dd>SI-ES3</dd>
  <dt>Vendor</dt><dd>Yaskawa Electric Corporation</dd>
  <dt>Documentation</dt><dd><a href=""></a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data) describes 1 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r37442</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td ><pre>SI-ES3</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x53455333</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x92420300</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=33 valign=top>TX PDOs</td>
<td><pre>0x1a00: Inputs</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a01: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6061:00  Modes of operation display      SINT (8 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a05: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6044:00  vl control effort               INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a06: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x60fd:00  Digital inputs                  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a14: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6042:00  vl target velocity              INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a15: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6043:00  vl velocity demand              INT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a16: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6048:01  vl velocity accel: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6048:02  vl velocity accel: delta time   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a17: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6049:01  vl velocity decel: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6049:02  vl velocity decel: delta time   UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a18: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x604a:01  vl velocity quick-stop: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x604a:02  vl velocity quick-stop: delta time  UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a19: Inputs</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x604c:01  vl dimension factor: numerator  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x604c:02  vl dimension factor: denumerator  DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a23: Inputs</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a24: Inputs</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a25: Inputs</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a26: Inputs</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a27: Inputs</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a28: Inputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=32 valign=top>RX PDOs</td>
<td><pre>0x1600: Outputs</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1601: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1605: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6042:00  vl target velocity              UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1606: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x60fe:01  Physical outputs                UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1607: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6060:00  Modes of operation              SINT (8 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1614: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6048:01  vl velocity accel: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6048:02  vl velocity accel: delta time   UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1615: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6049:01  vl velocity decel: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6049:02  vl velocity decel: delta time   UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1616: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x604a:01  vl velocity quick-stop: delta speed  UDINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x604a:02  vl velocity quick-stop: delta time  UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1617: Outputs</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x604c:01  vl dimension factor: numerator  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x604c:02  vl dimension factor: denumerator  DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1623: Outputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1624: Outputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1625: Outputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1626: Outputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1627: Outputs</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1628: Outputs</pre></td>
</tr>
</table>
## Generic XML Example
<pre class="xml">
&lt;slave idx="ADDRESS" type="generic" vid="0x00000539" pid="0x53455333" configPdos="true"&gt;
  &lt;syncManager idx="3" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
  &lt;syncManager idx="" dir="Dunno"&gt;
  &lt;/syncManager&gt;
&lt;/slave&gt;
</pre>
