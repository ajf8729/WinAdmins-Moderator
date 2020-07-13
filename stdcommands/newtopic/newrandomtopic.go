package newtopic

import (
	"math/rand"
)

func randomTopic() string {
	return topicList[rand.Intn(len(topicList))]
}

var topicList = []string{
	"Lets talk about why cats can't decide to go inside or outside.",
	"What type of weather are you having?",
	"What do you think the global impact of fruit rollups is.",
	"Who you think the best member of the FRIENDS cast is.",
	"How much water do you drink on a daily basis? Is it enough?",
	"What is the longest time you've been away from home",
	"Should pluto be a planet or nah?",
	"Whats your favorite technology conference? Why?",
	"Would you eat at a restaurant that was really dirty if the food was amazing.",
	"Who is your favorite entertainer (comedian musician etc) - and why?",
	"What's your favorite holiday",
	"What's your least favorite holiday",
	"Is epoch time useful or nah?",
	"Who in your life brings you the most joy?",
	"What're the best and worst things about the marketplace where you get your apps?",
	"What piece of technology is really frustrating to use?",
	"Do you prefer short hair or long hair on a guy/girl?",
	"What type of kid were you (e.g. spoiled, rebellious, well-behaved, quiet, obnoxious...)?",
	"What restaurant do you eat at most?",
	"What's your favorite number? Why?",
	"When was the last time you went to a movie theater?",
	"Which do you prefer? Books or movies?",
	"Are you a saver or a spender?",
	"Do you usually achieve the goals you set? Why or why not?",
	"What's the worst fast food restaurant?",
	"What would be your perfect weekend?",
	"What do you do when you're bored?",
	"What was the best period of your life so far? What do you think will be the best period of your entire life?",
	"Records, tapes, CDs, MP3s, streaming. Which did you grow up with? What is good and bad about each?",
	"Do you think people read more or fewer books now than 50 years ago?",
	"Are you very active, or do you prefer to just relax in your free time?",
	"What makes you nervous?",
	"Is cereal soup? Why or why not?",
	"What secret conspiracy would you like to start?",
	"What’s invisible but you wish people could see?",
	"What’s the weirdest smell you have ever smelled?",
	"Is a hotdog a sandwich? Why or why not?",
	"What’s the best Wi-Fi name you’ve seen?",
	"What’s the most ridiculous fact you know?",
	"What is something that everyone looks stupid doing?",
	"What is the funniest joke you know by heart?",
	"In 40 years, what will people be nostalgic for?",
	"What are the unwritten rules of where you work?",
	"How do you feel about putting pineapple on pizza?",
	"What part of a kid’s movie completely scarred you?",
	"What kind of secret society would you like to start?",
	"If animals could talk, which would be the rudest?",
	"Toilet paper, over or under?"
	"What’s the best type of cheese?",
	"What’s the best inside joke you’ve been a part of?",
	"In one sentence, how would you sum up the internet?",
	"How many chickens would it take to kill an elephant?",
	"What is the most embarrassing thing you have ever worn?",
	"What’s the weirdest thing a guest has done at your house?",
	"What mythical creature would improve the world most if it existed?",
	"What inanimate object do you wish you could eliminate from existence?",
	"What is the weirdest thing you have seen in someone else’s home?",
	"What would be the absolute worst name you could give your child?",
	"What are some of the nicknames you have for customers or coworkers?",
	"If peanut butter wasn’t called peanut butter, what would it be called?",
	"What movie would be greatly improved if it was made into a musical?",
	"What is the funniest corporate / business screw up you have heard of?",
	"What would be the worst “buy one get one free” sale of all time?",
	"If life were a video game, what would some of the cheat codes be?",
	"What is the funniest name you have actually heard used in the real world?",
	"What sport would be the funniest to add a mandatory amount of alcohol to?",
	"What would be the coolest animal to scale up to the size of a horse?",
	"What two totally normal things become really weird if you do them back to back?",
	"What is something that you just recently realized that you are embarrassed you didn’t realize earlier?",
	"What would be the best-worst name for different types of businesses? (dry cleaners, amusement parks, etc.)",
	"What are some things that are okay to occasionally do but definitely not okay to do every day?",
	"If you were arrested with no explanation, what would your friends and family assume you had done?",
	"Who would do you call if you get arrested?",
	"Where is the most beautiful place near where you live?",
	"What bands or types of music do you listen to when you exercise?",
	"What book has changed one of your long-held opinions?",
	"What is the most annoying thing about your phone?",
	"How often do you go to the library?",
	"When was the last time you had a food fight?",
	"Have you ever had a pillow fight?",
	"Did you own a wrist calculator?",
	"Intel or AMD?",
	"Linux or Windows?",
	"What's the best part of having kids? - if you have them.",
	"Have you ever contributed to an open source project?",
	"How much time to do you spend on WinAdmins every week?",
	"How much time do you spend doing 'work' things during non work time.",
	"Have you ever been called a nerd?",
	"What got you started in IT?",
	"What was your favorite IT project?",
	"What was the worst IT project you were ever involved in",
	"How much time off is just right?",
	"Have you ever taken more than two weeks of work off on purpose?",
	"If you had to pick one food to eat for th rest of your life what would it be?",
	"What was your favorite operating system - Its OK to say XP no one will judge you. Much.",
	"If you had a parrot what would you teach it to say?",
	"What is something you learned this week?",
	"What is somethign you learned in highschool?",
	"What is your favorite e-mail client for mobile?",
	"What is your dream retirement location?",
	"When do you think you'll retire?",	
}
