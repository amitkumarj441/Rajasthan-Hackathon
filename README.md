# Rajasthan-Hackathon
Developing {Manhood/Humanity}Coin based on Blockchain concept.

Description : Developing {Manhood/Humanity}Coins point system where users can honour each other with ManhoodCoins. Later ManhoodCoins can also be collected by wearable devices, or smart meters, rewarding people saving resources (electricity, water, gas) or living a healthy lifestyle reported by their wearable device. To avoid abuse of the system these points are audited by a tweeting auditor (using twitter). The ManhoodCoins points can be used by companies and government bodies to reward people doing good to their communities, health and to the environment.

## Technical Details

Types of thanks:
  - small = 1 {Manhood/Humanity}Coins
  - medium= 5 {Manhood/Humanity)Coins
  - large = 10{Manhood/Humanity}Coins
  
Attributes of a user:
  1. userID   (unique string, will be used as key)
  2. balance  (int, computed points from the type of thank)
  3. thanklist(string slice (array), array of the thanks received by the user)
  
Attributes of a thank:
  1. Thanker  (the name of the person giving the thank)
  2. ThankType(type of the thank small, medium, large)
  3. message  (a small message explaining the thank, can be empty)
  
