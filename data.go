package data

import (
	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/models"
)


type SeedData struct {
  Users  []models.User
  Products  []models.Product
  FeaturedProducts []models.FeaturedProduct
}

var InitialData SeedData = SeedData{
  Users: []models.User{
    {
      Name:     "John",
      Email:  "johnmwanzia277@gmail.com",
      Password: "987654321",
      IsAdmin:  true,
    },
    {
      Name:     "User1",
      Email:    "user@gmail.com",
      IsAdmin: false,
    },
  },
  Products: []models.Product{
    {
      Brand:     "Hp",
      CompName: "hp elitebook 840 g4 ",
      Slug:      "hp-elitebook-840-g4",
      Category:  "laptops",
      ImageUrl:   "https://cdn.shopify.com/s/files/1/1918/7539/products/Hp-840-g4-touch.png?v=1647739774",
      Price:     35000,
      CountInStock: 4,
      Specs:    " Intel Core i7-6300U 2.60GHz 8GB DDR4 RAM 256GB SSD 14 LED Display",
    },
    {
      Brand:     "Hp",
      CompName: "Hp Elitebook 840 G3",
      Slug:      "Hp-Elitebook-840-G3",
      Category:  "laptops",
      ImageUrl:   "https://w7.pngwing.com/pngs/548/236/png-transparent-hp-elitebook-840-g3-laptop-hp-elitebook-820-g3-hp-elitebook-745-g3-laptop-electronics-netbook-computer.png",
      Price:     35000,
      CountInStock: 7,
      Specs:    " Intel Core i7-6300U 2.60GHz 8GB DDR4 RAM 256GB SSD 14 LED Display",

    },
    {
      Brand:     "lenovo",
      CompName: "Lenovo ThinkPad T470s ",
      Slug:     "Lenovo-ThinkPad-T470s",
      Category:  "laptops",
      ImageUrl:   "https://www.lenovo.com/medias/lenovo-laptop-thinkpad-t470s-hero.png?context=bWFzdGVyfGltYWdlc3wyMjU4NjF8aW1hZ2UvcG5nfGltYWdlcy9oN2YvaDAzLzE0MzMyNjg1MDkwODQ2LnBuZ3w2ZWY5N2ZjZjQ3OGY1MTFkNTlkMmI2YjY3NjZmNDEyOGQxZTFiYWVmNzgwNjY2NTJjZDRlYzk1MTRmZjU0MjVl",
      Price:     35000,
      CountInStock: 6,
      Specs:    " Intel Core i7-6300U 2.60GHz 8GB DDR4 RAM 256GB SSD 14 LED Display",
    },
    {
      Brand:     "Apple",
      CompName: "Macbook Pro 16  MVVK2",
      Slug:     "Macbook-Pro-16-MVVK2",
      Category:  "laptops",
      ImageUrl:   "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTe2dwWjlnZOSiKq03vu6k0iSCYzO0kp9CNgg&usqp=CAU",
      Price:     35000,
      CountInStock: 3,
      Specs:    " Intel Core i7-6300U 2.60GHz 8GB DDR4 RAM 256GB SSD 14 LED Display",
    },
  },
  FeaturedProducts: []models.FeaturedProduct{
    {
      GadgetName: " samsung-a23 ",
      Slug:        "samsung-a23",
      Category:    "phones",
      ImageUrl:     "https://www.91-cdn.com/hub/wp-content/uploads/2022/02/Samsung-Galaxy-A23-5G-6.jpeg",
      Price:       35000,
      CountInStock: 4,
      Specs:       " 6.5-inch HD+ Infinity-V display, 1080 x 2408 pixels, 20:9 aspect ratio, 90Hz refresh rate, 270 PPI Android 12 	64GB 4GB RAM,",
    },
    {
      GadgetName: " huawei honor 10  ",
      Slug:       "huawei-honor-10 ",
      Category:   "phones",
      ImageUrl:    "https://www.priceinkenya.com/_ipx/b_%23ffffff,f_png,q_75,fit_contain,s_300x300/https://api.priceinkenya.com/media/121038/conversions/huawei-honor-10-64gb-zwXAQG9nI2-original.webp",
      Price:      35000,
      CountInStock: 4,
      Specs:      " 6.5-inch HD+ Infinity-V display, 1080 x 2408 pixels, 20:9 aspect ratio, 90Hz refresh rate, 270 PPI Android 12 	64GB 4GB RAM,",
    },
    {
      GadgetName: "headphones",
      Slug:        "headphones-1",
      Category:    "headphones",
      ImageUrl:     "https://t4.ftcdn.net/jpg/05/35/49/97/240_F_535499727_fLDl5BxuoURnqRVLJ9jc46rBMws9lgMs.jpg",
      Price:      1500,
      CountInStock: 4,
      Specs:      "Sennheiser Pro Audio Model	HD 280 Pro white	Wireless",
    },
    {
      GadgetName: "hp laptop charger",
      Slug:        "hp-laptop-charger",
      Category:    "chargers",
      ImageUrl:     "https://3.imimg.com/data3/WG/AR/MY-22059447/19-5v-4-62a-adapter-500x500.jpg",
      Price:     2500,
      CountInStock: 4,
      Specs:     "HP 19.5V 4.62A 90W Laptop Charger AC/DC Adapter",
    },
    {
      GadgetName: "iphone 14",
      Slug:        "iphone-14",
      Category:    "phones",
      ImageUrl:     "https://s7d1.scene7.com/is/image/dish/iPhone_14_Pro_Max_Deep_Purple_phonewall?$ProductBase$",
      Price:    95000,
      CountInStock: 4,
      Specs:    " iPhone 14 Pro Max 256GB 6GB RAM 6.7 inches 5G 20MP Camera 4K Video 5000mAh Battery ",
    },
    {
      GadgetName: "wireless mouse",
      Slug:        "wireless-mouse",
      Category:    "mouse",
      ImageUrl:     "https://atlas-content-cdn.pixelsquid.com/assets_v2/200/2009721087641786232/previews/G03-200x200.jpg",
      Price:      1000,
      CountInStock: 4,
      Specs:      "Wireless Mouse for Laptop, Inphic 2.4G Rechargeable Silent Computer Mouse,1600 DPI Ultra Thin Optical Portable USB Mini Mouse, Cordless Mice for Laptop,PC,MacBook,Mac, white",
    },
    {
      GadgetName: "webcam",
      Slug:        "webcam",
      Category:    "webcam",
      ImageUrl:     "https://w7.pngwing.com/pngs/302/546/png-transparent-webcam-camera-microphone-internet-webcam-electronics-microphone-camera-lens-thumbnail.png",
      Price:     5000,
      CountInStock: 4,
      Specs:     "Webcam with Microphone, 1080P HD Streaming USB Computer Webcam [Plug and Play] [30fps] ",
    },
    {
      GadgetName: "wireless keyboard",
      Slug:        "wireless-keyboard",
      Category:    "keyboard", 
      ImageUrl:   "https://png.pngtree.com/png-vector/20210205/ourlarge/pngtree-wireless-keyboard-png-image_2872698.jpg",
      Price:     1500,
      CountInStock: 4,
      Specs:     " Jelly Comb 2.4G Slim Ergonomic Quiet Keyboard ",
    },

},

}
