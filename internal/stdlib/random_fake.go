package stdlib

import (
  "github.com/gammazero/deque"
  "github.com/brianvoe/gofakeit/v6"
  tc "github.com/vulogov/ThreadComputation"
)

func BUNDrandomName(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.Name(), nil
}

func BUNDrandomEmail(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.Email(), nil
}

func BUNDrandomPhone(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.PhoneFormatted(), nil
}

func BUNDrandomCompany(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.Company(), nil
}

func BUNDrandomCC(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.CreditCardNumber(nil), nil
}

func BUNDrandomSSN(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.SSN(), nil
}

func BUNDrandomPassword(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.Password(true, true, true, true, false, 16), nil
}

func BUNDrandomHipster(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.HipsterWord(), nil
}

func BUNDrandomAddress(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  a := gofakeit.Address()
  res := tc.MakeDict()
  res.D.Set("address", a.Address)
  res.D.Set("street", a.Street)
  res.D.Set("city", a.City)
  res.D.Set("state", a.State)
  res.D.Set("zip", a.Zip)
  res.D.Set("country", a.Country)
  res.D.Set("lat", a.Latitude)
  res.D.Set("lon", a.Longitude)
  return res, nil
}

func BUNDrandomTextParagraph(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  return gofakeit.Paragraph(1, 25, 125, " "), nil
}

func init() {
  faker := gofakeit.NewCrypto()
  gofakeit.SetGlobalFaker(faker)
  tc.SetCommand("random.Name", BUNDrandomName)
  tc.SetCommand("random.Email", BUNDrandomEmail)
  tc.SetCommand("random.Phone", BUNDrandomPhone)
  tc.SetCommand("random.Company", BUNDrandomCompany)
  tc.SetCommand("random.CreditCard", BUNDrandomCC)
  tc.SetCommand("random.SSN", BUNDrandomSSN)
  tc.SetCommand("random.Password", BUNDrandomPassword)
  tc.SetCommand("random.Hipster", BUNDrandomHipster)
  tc.SetCommand("random.Address", BUNDrandomAddress)
  tc.SetCommand("random.Text", BUNDrandomTextParagraph)
}
