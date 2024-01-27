<div class="nav"><a href="/esi-data">Home</a> | <a href="/esi-data/devices">Device List</a> | Unidrive M702 RFC-S</div>

#  Control Unidrive M702 RFC-S

<dl>
  <dt>Type:</dt><dd>Unidrive M702 RFC-S</dd>
  <dt>Description:</dt><dd>Unidrive M702 RFC-S</dd>
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
<td ><pre>Unidrive M702 RFC-S</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td ><pre>0x01030302</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0200000c</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td ><pre><a href="Digitax+HD+M750+RFC-A">Digitax HD M750 RFC-A r512</a><br/><a href="Digitax+HD+M750+RFC-S">Digitax HD M750 RFC-S r512</a><br/><a href="Digitax+HD+M751+RFC-A">Digitax HD M751 RFC-A r512</a><br/><a href="Digitax+HD+M751+RFC-S">Digitax HD M751 RFC-S r512</a><br/><a href="Digitax+HD+M753+RFC-A">Digitax HD M753 RFC-A r512</a><br/><a href="Digitax+HD+M753+RFC-S">Digitax HD M753 RFC-S r512</a><br/><a href="Unidrive+HS70+RFC-A">Unidrive HS70 RFC-A r512</a><br/><a href="Unidrive+HS70+RFC-S">Unidrive HS70 RFC-S r512</a><br/><a href="Unidrive+HS71+RFC-A">Unidrive HS71 RFC-A r512</a><br/><a href="Unidrive+HS71+RFC-S">Unidrive HS71 RFC-S r512</a><br/><a href="Unidrive+HS72+RFC-A">Unidrive HS72 RFC-A r512</a><br/><a href="Unidrive+HS72+RFC-S">Unidrive HS72 RFC-S r512</a><br/><a href="Unidrive+M600+RFC-A">Unidrive M600 RFC-A r512</a><br/><a href="Unidrive+M600+RFC-S">Unidrive M600 RFC-S r512</a><br/><a href="Unidrive+M700+RFC-A">Unidrive M700 RFC-A r512</a><br/><a href="Unidrive+M700+RFC-S">Unidrive M700 RFC-S r512</a><br/><a href="Unidrive+M701+RFC-A">Unidrive M701 RFC-A r512</a><br/><a href="Unidrive+M701+RFC-S">Unidrive M701 RFC-S r512</a><br/><a href="Unidrive+M702+RFC-A">Unidrive M702 RFC-A r512</a><br/><a href="Unidrive+M880+RFC-A">Unidrive M880 RFC-A r512</a><br/><a href="Unidrive+M880+RFC-S">Unidrive M880 RFC-S r512</a><br/><a href="Unidrive+M881+RFC-A">Unidrive M881 RFC-A r512</a><br/><a href="Unidrive+M881+RFC-S">Unidrive M881 RFC-S r512</a><br/><a href="Unidrive+M882+RFC-A">Unidrive M882 RFC-A r512</a><br/><a href="Unidrive+M882+RFC-S">Unidrive M882 RFC-S r512</a><br/><a href="Unidrive+M888+RFC-A">Unidrive M888 RFC-A r512</a><br/><a href="Unidrive+M888+RFC-S">Unidrive M888 RFC-S r512</a><br/><a href="Unidrive+M889+RFC-A">Unidrive M889 RFC-A r512</a><br/><a href="Unidrive+M889+RFC-S">Unidrive M889 RFC-S r512</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=14 valign=top>TX PDOs</td>
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
<td ><pre>0x1a02: Transmit PDO Mapping 3</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6064:00  Position actual value           DINT (32 bits)</pre></td>
</tr>
<tr class="txpdo pdosection">
<td ><pre>0x1a04: Transmit PDO Mapping 5</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6041:00  Status word                     UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td ><pre>  0x6077:00  Torque actual value             INT (16 bits)</pre></td>
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
<td class="first" rowspan=14 valign=top>RX PDOs</td>
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
<td ><pre>0x1602: Receive PDO Mapping 3</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x607a:00  Target position                 DINT (32 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td ><pre>0x1604: Receive PDO Mapping 5</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6040:00  Control word                    UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td ><pre>  0x6071:00  Target torque                   INT (16 bits)</pre></td>
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
