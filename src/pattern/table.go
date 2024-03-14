package pattern

const IndexTmpl = `<!DOCTYPE html>
<html>
<style>
table, th, td {
  border:1px solid black;
}
</style>
<body>
<table style="width:100%">

  <th colspan="2">Order</th>
  <tr>
    <td>order_uid</td>
    <td>{{ .OrderUid }}</td>
  </tr>
  <tr>
    <td>track_number</td>
    <td>{{ .TrackNumber  }}</td>
  </tr>
  <tr>
    <td>entry</td>
    <td>{{ .Entry }}</td>
  </tr>
  <tr>
    <td>delivery</td>
    <td>Name: {{ .Delivery.Name }} <br> 
    Phone: {{ .Delivery.Phone }}<br> 
    Zip: {{ .Delivery.Zip }}<br>
    City: {{ .Delivery.City }}<br>
    Address: {{ .Delivery.Address }}<br>
    Region: {{ .Delivery.Region }}<br>
    Email: {{ .Delivery.Email }}<br></td>
  </tr>
  <tr>
    <td>payment</td>
    <td>Transaction: {{ .Payment.Transaction }} <br> 
    RequestId: {{ .Payment.RequestId }}<br> 
    Currency: {{ .Payment.Currency }}<br>
    Provider: {{ .Payment.Provider }}<br>
    Amount: {{ .Payment.Amount }}<br>
    PaymentDt: {{ .Payment.PaymentDt }}<br>
    Bank: {{ .Payment.Bank }}<br>
    DeliveryCost: {{ .Payment.DeliveryCost }}<br>
    GoodsTotal: {{ .Payment.GoodsTotal }}<br>
    CustomFee: {{ .Payment.CustomFee }}<br>
	</td>
  </tr>
  <tr>
    <td>items</td>
    <td>{{range .Items}}
    ChrtId: {{ .ChrtId }}<br> 
    TrackNumber: {{ .TrackNumber }}<br>
    Price: {{ .Price }}<br>
    Rid: {{ .Rid }}<br>
    Name: {{ .Name }}<br>
    Sale: {{ .Sale }}<br>
    Size: {{ .Size }}<br>
    TotalPrice: {{ .TotalPrice }}<br>
    NmId: {{ .NmId }}<br>
    Brand: {{ .Brand }}<br>
    Status: {{ .Status }}<br>
<hr>
	  {{end}}</td>
  </tr>
  <tr>
    <td>locale</td>
    <td>{{ .Locale }}</td>
  </tr>
  <tr>
    <td>internal_signature</td>
    <td>{{ .InternalSignature }}</td>
  </tr>
  <tr>
    <td>customer_id</td>
    <td>{{ .CustomerId }}</td>
  </tr>
  <tr>
    <td>delivery_service</td>
    <td>{{ .DeliveryService }}</td>
  </tr>
  <tr>
    <td>shardkey</td>
    <td>{{ .Shardkey }}</td>
  </tr>
  <tr>
    <td>sm_id</td>
    <td>{{ .SmId }}</td>
  </tr>
  <tr>
    <td>date_created</td>
    <td>{{ .DateCreated }}</td>
  </tr>
  <tr>
    <td>oof_shard</td>
    <td>{{ .OofShard }}</td>
  </tr>
</table>
</body>
</html>
`
