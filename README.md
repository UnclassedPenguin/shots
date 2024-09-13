## Shots Tracker

Hmm..


//ID 
//Date &timeStr
//Gun (-g string) &gun

Ammo Type (-at string) &ammoType
Ammo Weight (-aw int) &ammoWeight

//Number of Shots (-n int)

Ammo Price (-ap float64) &ammoIndivPrice
Total Price (float64) &ammoTotalPrice

//Notes

d.AddRecord(db, timeStr, gun, ammoType, ammoWeight, number, ammoIndivPrice, AmmoTotalPrice, notes)

### To-do
- Add a projectile weight category
- Add a projectile type category
- Add a projectile cost category
- Add a total cost for each row
