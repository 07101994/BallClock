using System;

namespace BallClockDotNet
{
	class Program
	{
		static void Main(string[] args)
		{
			int size = 0;

			while (int.TryParse(Console.ReadLine(), out size))
			{
				if (size == 0) break;
				if (size >= 27 && size <= 127)
				{
					var clock = new Clock(size);
					int result = clock.GetLCM();
					Console.WriteLine("{0} balls cycle after {1} days.", size, result / 2);
				}
			}
		}
	}
}
