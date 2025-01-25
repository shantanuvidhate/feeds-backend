package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"github.com/shantanuvidhate/feeds-backend/internal/store"
)

var usernames = []string{
	"CoolCucumber", "TechWizard", "FunkyMonkey", "PixelHunter", "SwiftShadow",
	"CyberNinja", "BlazePhoenix", "MysticVoyager", "QuantumFrost", "DigitalDreamer",
	"NeonKnight", "SonicTurtle", "EpicVoyager", "TurboPanda", "MagicFalcon",
	"StarGazer", "FrostyBear", "DragonSlayer", "CodeGuru", "StealthTiger",
	"LunarWolf", "ThunderBolt", "CryptoWizard", "IronLion", "CosmicJester",
	"VelocityViper", "ShadowSniper", "SkyDiver", "MegaPixel", "FireStorm",
	"SilentDagger", "IcePhantom", "QuantumPilot", "DreamCatcher", "GhostRider",
	"NovaBlaster", "SolarEagle", "SteelSamurai", "AlphaPredator", "NightHawk",
	"TurboRocket", "BlizzardRider", "GalaxyNomad", "ElectricWolf", "AstroRanger",
	"MysticDancer", "PhoenixBlaze", "ShadowHunter", "IronFalcon", "SkyRanger",
	"StealthBlade", "MegaFox", "DigitalLion", "EpicVortex", "ThunderEagle",
	"LunarSniper", "BlazeNomad", "CodeMaster", "MysticShark", "FrozenWolf",
	"CosmicVoyager", "ElectricPanda", "AstroNomad", "MagicWolf", "FirePhoenix",
	"SilentFox", "SkySamurai", "TurboLion", "BlizzardNomad", "GalaxyWolf",
	"IronRider", "SolarPhoenix", "ShadowTiger", "QuantumSamurai", "FrostBlaster",
	"DreamNomad", "GhostBlaze", "NovaHunter", "StarPilot", "VelocityRocket",
	"AlphaWolf", "ElectricEagle", "AstroTiger", "ThunderWolf", "LunarEagle",
	"MagicNomad", "IceRider", "QuantumFox", "BlazeHunter", "CosmicFalcon",
	"FrozenNomad", "StealthRanger", "TurboNomad", "SkyTiger", "MegaEagle",
	"DigitalSamurai", "EpicWolf", "ThunderTiger", "MysticEagle", "FireBlaster",
}

var titles = []string{
	"The Art of Minimalism: Simplify Your Life",
	"Exploring the Cosmos: A Beginner’s Guide to Astronomy",
	"Mastering Productivity: Hacks for a Busy Life",
	"The Future of AI: Opportunities and Challenges",
	"Cooking with Love: Easy Recipes for Everyone",
	"The Hidden Wonders of the Ocean",
	"From Hobby to Hustle: Turning Passion into Profit",
	"A Journey Through Time: The History of Ancient Civilizations",
	"The Psychology of Happiness: What Science Says",
	"Building Your Dream Home on a Budget",
	"Top 10 Travel Destinations for Adventure Seekers",
	"How to Start Your Own Business: A Step-by-Step Guide",
	"The Science Behind Climate Change",
	"The Secret to Successful Relationships",
	"Fitness for Beginners: Your 30-Day Plan",
	"Mindful Living: Tips for a Balanced Life",
	"Exploring the World of Virtual Reality",
	"The Power of Music: How It Shapes Our Lives",
	"Understanding Blockchain Technology",
	"Gardening for Beginners: Tips to Grow Your Green Thumb",
	"The Ultimate Guide to Freelancing",
	"Top 5 Trends in Web Development for 2025",
	"The Healing Power of Nature: Why You Should Spend Time Outdoors",
	"The Beginner’s Guide to Cryptocurrency",
	"10 Books That Will Change Your Perspective on Life",
	"The Evolution of Technology: Past, Present, and Future",
	"DIY Home Projects to Transform Your Space",
	"How to Create a Winning Resume",
	"The Magic of Storytelling: How to Captivate an Audience",
	"Sustainable Living: Small Steps to Make a Big Impact",
	"The Art of Negotiation: Tips for Success",
	"Discovering Hidden Gems: Underrated Travel Destinations",
	"Mastering the Digital Workspace: Tools and Tips",
	"How to Build a Strong Personal Brand",
	"The Role of Artificial Intelligence in Healthcare",
	"Top 10 Strategies for Effective Leadership",
	"Understanding the Basics of Stock Market Investing",
	"The Power of Habit: Transform Your Life in 30 Days",
	"Exploring the Science of Sleep: Why It Matters",
	"The Beginner’s Guide to Photography",
	"Top 5 Mistakes to Avoid When Starting a Business",
	"The History of Art: A Visual Journey Through Time",
	"Cybersecurity 101: Protecting Your Digital Life",
	"How to Stay Motivated and Achieve Your Goals",
	"The Science of Meditation: Benefits for Mind and Body",
	"How to Plan the Perfect Road Trip",
	"The Future of Renewable Energy: What’s Next?",
	"How to Improve Your Communication Skills",
	"The Rise of E-Sports: A New Era in Gaming",
	"From Zero to Hero: Inspiring Stories of Success",
}

var contents = []string{
	"The rapid advancement of technology has brought both opportunities and challenges to modern society. Artificial intelligence, for example, is revolutionizing industries by automating processes and analyzing data at unprecedented speeds. However, it also raises concerns about job displacement and ethical dilemmas. Balancing innovation with responsibility is crucial. Education systems must adapt to equip individuals with the skills needed to thrive in this changing landscape. Furthermore, fostering collaboration between governments, businesses, and communities is essential to ensure that technology benefits everyone. As we navigate this digital era, embracing both progress and ethical considerations will pave the way for a more equitable future.",
	"Travel is more than just a physical journey; it’s an opportunity to experience new cultures, meet people from diverse backgrounds, and broaden your perspective on life. Whether you’re exploring ancient ruins, savoring exotic cuisines, or trekking through breathtaking landscapes, each adventure leaves you with memories that last a lifetime. Travel also challenges you to step out of your comfort zone, adapt to new environments, and develop a deeper appreciation for the world’s diversity. It’s a powerful way to gain insight into different ways of living and thinking. Ultimately, travel enriches your soul and fosters a greater sense of connection.",
	"Remote work has become a defining feature of the modern professional landscape, offering flexibility and convenience to employees across industries. With advancements in communication tools and project management software, teams can collaborate seamlessly, regardless of their physical location. However, remote work also presents unique challenges, such as maintaining work-life balance, combating isolation, and fostering team cohesion. Employers must prioritize creating inclusive virtual environments and providing resources to support their teams. For employees, developing strong self-discipline and effective communication skills is key. As the world embraces this shift, remote work has the potential to redefine productivity and unlock new possibilities.",
	"Climate change represents one of the most pressing challenges of our time. Rising global temperatures, melting polar ice caps, and increasingly severe weather events highlight the urgent need for action. Addressing climate change requires a collective effort from individuals, governments, and industries. Renewable energy sources, such as solar and wind power, offer viable alternatives to fossil fuels, while sustainable practices like reducing waste and conserving resources play a vital role. Education and awareness are equally important in driving change. By adopting eco-friendly habits and supporting policies that prioritize environmental preservation, we can work together to safeguard the planet.",
	"Healthy living encompasses more than just physical fitness; it’s about achieving balance in all aspects of life. Nutrition, exercise, sleep, and mental well-being are interconnected components that contribute to overall health. A balanced diet rich in whole foods provides essential nutrients, while regular physical activity strengthens the body and boosts energy levels. Adequate sleep allows the body to recover and rejuvenate, while mindfulness practices like meditation promote mental clarity and reduce stress. By prioritizing these areas, individuals can enhance their quality of life and build resilience against illness. Healthy living is a lifelong journey that fosters vitality and fulfillment.",
	"The history of human civilization is a testament to our ingenuity, resilience, and desire for progress. From the earliest hunter-gatherer societies to the development of agriculture, humans have continually adapted to their environments. Ancient civilizations, such as Mesopotamia, Egypt, and the Indus Valley, laid the foundation for modern society with their innovations in writing, architecture, and governance. Over centuries, cultural exchange, exploration, and technological advancements have shaped the world we live in today. Studying history allows us to understand our shared heritage, learn from past mistakes, and appreciate the interconnectedness of humanity. It’s a story of triumph and transformation.",
	"Learning new skills is one of the most empowering experiences an individual can undertake. Whether it’s mastering a musical instrument, coding a program, or developing public speaking abilities, the process of learning stimulates the mind and fosters personal growth. It requires dedication, practice, and a willingness to embrace challenges. In today’s fast-paced world, the ability to adapt and acquire new skills is more valuable than ever. Continuous learning not only enhances career prospects but also enriches life experiences. It instills a sense of accomplishment and opens doors to new opportunities, proving that growth is a lifelong journey worth pursuing.",
	"Storytelling is an ancient art form that has transcended generations and cultures. From oral traditions to written literature and modern media, stories have the power to captivate, inspire, and educate. They allow us to explore complex emotions, share experiences, and convey universal truths. A compelling story often includes relatable characters, a clear narrative arc, and a message that resonates with the audience. Whether through books, films, or digital platforms, storytelling connects people on a profound level, fostering empathy and understanding. In a world filled with diverse perspectives, the ability to tell and appreciate stories remains a timeless human skill.",
	"The digital age has revolutionized the way we access and share information. With the internet, knowledge is at our fingertips, enabling us to learn about any topic, connect with others globally, and stay informed about current events. However, this convenience comes with challenges, such as misinformation and digital fatigue. To navigate the digital landscape effectively, it’s essential to develop critical thinking skills and establish healthy boundaries. Embracing technology responsibly can enhance productivity, foster creativity, and improve communication. As we continue to integrate digital tools into our lives, finding a balance between connectivity and mindfulness is key to long-term well-being.",
	"Renewable energy is a cornerstone of the global effort to combat climate change and reduce reliance on fossil fuels. Solar panels, wind turbines, and hydroelectric dams are becoming increasingly efficient and accessible, offering sustainable alternatives for power generation. Transitioning to renewable energy not only reduces greenhouse gas emissions but also creates jobs and stimulates economic growth. Governments, businesses, and communities must work together to invest in clean energy infrastructure and research. Public awareness and support are crucial for driving this transition. By embracing renewable energy, we can build a greener, more sustainable future for generations to come.",
	"Mental health is an integral part of overall well-being, yet it has often been overlooked or stigmatized. In recent years, there has been a growing awareness of the importance of mental health, leading to more open conversations and improved access to resources. Practices such as mindfulness, therapy, and self-care are becoming mainstream, helping individuals manage stress and build resilience. Employers are recognizing the value of supporting mental health in the workplace, creating environments where employees can thrive. By continuing to prioritize mental health, society can foster a culture of empathy and understanding, ultimately enhancing the quality of life for all.",
	"Space exploration represents humanity’s enduring curiosity and desire to push boundaries. From the first moon landing to the development of Mars rovers, each milestone demonstrates our ability to overcome challenges and innovate. Space missions contribute to scientific knowledge, technological advancements, and a deeper understanding of our universe. They also inspire future generations to pursue careers in science, technology, engineering, and mathematics (STEM). As private companies join national space agencies in exploring the cosmos, collaboration and sustainable practices will be essential. The dream of reaching new worlds and uncovering the mysteries of the universe continues to unite and inspire humanity.",
	"Entrepreneurship is a dynamic and rewarding journey that requires vision, resilience, and adaptability. Starting a business involves identifying a problem, creating a solution, and bringing it to market. Entrepreneurs face numerous challenges, including securing funding, building a team, and navigating competition. However, the rewards—both personal and professional—can be immense. Successful entrepreneurs often share common traits, such as perseverance, creativity, and a willingness to take risks. In today’s rapidly changing world, entrepreneurship plays a vital role in driving innovation and economic growth. By fostering an entrepreneurial mindset, individuals can create opportunities, solve problems, and make a meaningful impact on society.",
	"Education is the foundation upon which societies are built. It empowers individuals with knowledge, critical thinking skills, and the ability to contribute meaningfully to their communities. Access to quality education remains a global challenge, with millions of children lacking basic learning resources. Investing in education, particularly for marginalized groups, has a ripple effect, improving health, economic stability, and social equality. Modern education systems are also evolving, incorporating technology and innovative teaching methods to prepare students for the future. By prioritizing education, we can unlock human potential and create a world where everyone has the opportunity to succeed.",
	"Culinary arts are a celebration of culture, creativity, and craftsmanship. From the careful selection of ingredients to the artistry of presentation, cooking is both a science and an art. It brings people together, fostering connection and joy through shared meals. Every dish tells a story, reflecting traditions, innovations, and personal expression. Whether it’s experimenting with fusion cuisine or preserving age-old recipes, culinary exploration enriches our understanding of diverse cultures. Food has the power to comfort, nourish, and delight, making it a universal language that transcends borders. Embracing the culinary arts is an invitation to savor the beauty of life.",
	"Fitness is a holistic pursuit that encompasses physical, mental, and emotional well-being. Regular exercise strengthens the body, improves cardiovascular health, and boosts energy levels. It also has profound effects on mental health, reducing stress and enhancing mood through the release of endorphins. Beyond physical activity, fitness involves proper nutrition, hydration, and rest to support overall health. In a fast-paced world, finding time for fitness can be challenging, but the benefits are undeniable. Whether through structured workouts, outdoor activities, or mindfulness practices like yoga, fitness empowers individuals to live vibrant and fulfilling lives. It’s an investment in longevity and vitality.",
	"The evolution of technology has reshaped industries and transformed everyday life. From the advent of the internet to the rise of artificial intelligence, technological innovations have revolutionized communication, healthcare, transportation, and more. These advancements have made the world more connected and efficient, but they also pose challenges, such as data privacy and ethical considerations. As technology continues to evolve, it’s essential to strike a balance between innovation and responsibility. By leveraging technology for the greater good, we can address global challenges, improve quality of life, and create a future where progress benefits everyone.",
	"Environmental conservation is critical to preserving the Earth’s natural resources and biodiversity. Deforestation, pollution, and habitat destruction threaten ecosystems and the species that depend on them. Conservation efforts, such as reforestation, wildlife protection, and sustainable practices, play a vital role in mitigating these impacts. Education and community involvement are key to driving change, empowering individuals to take action and advocate for the planet. By prioritizing conservation, we can ensure a healthier environment for future generations while maintaining the delicate balance of nature. It’s a collective responsibility that requires commitment and collaboration at every level of society.",
	"Art is a universal language that transcends cultural and linguistic barriers. It allows individuals to express emotions, ideas, and perspectives in ways that words cannot. From painting and sculpture to music and dance, art enriches our lives and inspires creativity. It also serves as a reflection of society, capturing historical moments and challenging norms. Supporting the arts is essential to fostering innovation and preserving cultural heritage. Whether through creating, appreciating, or sharing art, we can connect with others on a profound level. In a world often divided by differences, art has the power to unite and uplift humanity.",
	"Innovation is the driving force behind progress and transformation. It involves thinking outside the box, challenging conventional approaches, and developing new solutions to complex problems. Innovations in science, technology, and design have revolutionized industries and improved quality of life. From medical breakthroughs to advancements in renewable energy, innovation addresses pressing global challenges and creates opportunities for growth. Encouraging a culture of curiosity and experimentation is essential to fostering innovation. By embracing change and pushing boundaries, individuals and organizations can unlock their potential and contribute to a better future. Innovation is not just about ideas; it’s about making them a reality.",
}

var tags = []string{
	"technology",
	"travel",
	"health",
	"fitness",
	"education",
	"climatechange",
	"artificialintelligence",
	"renewableenergy",
	"mentalhealth",
	"entrepreneurship",
	"spaceexploration",
	"storytelling",
	"innovation",
	"culture",
	"history",
	"food",
	"nature",
	"leadership",
	"productivity",
	"sustainability",
}

var comments = []string{
	"This article was incredibly insightful and provided a fresh perspective on the topic!",
	"I found this content very helpful. Looking forward to more posts like this.",
	"Great points raised here. I especially liked the section about renewable energy.",
	"This was an enjoyable read! Thanks for sharing your thoughts.",
	"I disagree with some points, but overall, this was a well-written piece.",
	"Could you elaborate more on the challenges mentioned? That part was fascinating.",
	"This is exactly what I needed to read today. Thank you!",
	"The examples provided here were very relatable and practical.",
	"I appreciate the depth of research that went into this content.",
	"This raises some important questions. I’d love to hear others’ opinions on this.",
	"The writing style made this easy to follow and engaging. Well done!",
	"I’m inspired to take action after reading this. Thanks for the motivation!",
	"Some great takeaways here. I’ve bookmarked this for future reference.",
	"Can you provide additional resources or links for further reading?",
	"This touched on so many important issues. Thanks for highlighting them.",
	"The tone of this piece was very refreshing. Keep up the good work!",
	"I hadn’t thought about it this way before. This was an eye-opener!",
	"This is a topic that deserves more attention. Thanks for bringing it up.",
	"I really enjoyed the personal anecdotes included here—they made the piece relatable.",
	"There’s a lot to unpack here. I’ll be revisiting this post for sure!",
}

func Seed(store *store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.User.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Post.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comment.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return
		}
	}

	log.Panicln("Seeding completed successfully")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i), // Add a number to make it unique
			Email:    usernames[i] + "@email.com",
			Password: "password@123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))] // Randomly select a user

		posts[i] = &store.Post{
			UserId:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags:    []string{tags[rand.Intn(len(tags))], tags[rand.Intn(len(tags))]},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {

	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostId:  posts[rand.Intn(len(posts))].ID,
			UserId:  users[rand.Intn(len(users))].ID,
			Content: comments[rand.Intn(len(comments))],
		}
	}

	return cms
}
