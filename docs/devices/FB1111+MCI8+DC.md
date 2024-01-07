#  Beckhoff FB1111 MCI8 DC

<dl>
  <dt>Type:</dt><dd>FB1111 MCI8 DC</dd>
  <dt>Description:</dt><dd>FB1111 MCI8 DC 2xMII</dd>
  <dt>Vendor</dt><dd>Beckhoff Automation GmbH & Co. KG</dd>
  <dt>Documentation</dt><dd><a href="http://www.beckhoff.com/FB1111">http://www.beckhoff.com/FB1111</a></dd>
</dl>
## Revisions and PDOs
The ESI data ingested by [github.com/linuxcnc-ethercat/esi-data](http://github.com/linuxcnc-ethercat/esi-data)describes 2 revision(s) of this hardware.  Here are the known revisions and their differences.

This also includes the send and receive PDOs defined for each revision, and a link to other known devices with identical PDOs.

<table>
<tr >
<td class="first">Revision</td>
<td ><pre>r800</pre></td>
<td ><pre>r801</pre></td>
</tr>
<tr >
<td class="first">Name</td>
<td  colspan=2 align="center"><pre>FB1111 MCI8 DC 2xMII</pre></td>
</tr>
<tr >
<td class="first">PID</td>
<td  colspan=2 align="center"><pre>0x04570862</pre></td>
</tr>
<tr >
<td class="first">Revision Code</td>
<td ><pre>0x0320008c</pre></td>
<td ><pre>0x0321008c</pre></td>
</tr>
<tr >
<td class="first">Equivalant Devices</td>
<td  colspan=2 align="center"><pre><a href="FB1111+MCI16+DC">FB1111 MCI16 DC r1000</a><br/><a href="FB1111+MCI16+DC">FB1111 MCI16 DC r1001</a><br/><a href="FB1311+MCI16+DC">FB1311 MCI16 DC r1000</a><br/><a href="FB1311+MCI8+DC">FB1311 MCI8 DC r800</a></pre></td>
</tr>
<tr class="txpdo pdosection">
<td class="first" rowspan=3 valign=top>TX PDOs</td>
<td colspan=2 align="left"><pre>0x1600: Inputs</pre></td>
<td></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6000:01  Test In1                        UINT (16 bits)</pre></td>
</tr>
<tr class="txpdo">
<td  colspan=2 align="left"><pre>  0x6000:02  Test In2                        UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo pdosection">
<td class="first" rowspan=3 valign=top>RX PDOs</td>
<td colspan=2 align="left"><pre>0x1a00: Outputs</pre></td>
<td></td>
</tr>
<tr class="rxpdo">
<td  colspan=2 align="left"><pre>  0x7000:01  Test Out1                       UINT (16 bits)</pre></td>
</tr>
<tr class="rxpdo">
<td  colspan=2 align="left"><pre>  0x7000:02  Test Out2                       UINT (16 bits)</pre></td>
</tr>
</table>
