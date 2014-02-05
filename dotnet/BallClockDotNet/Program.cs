using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BallClockDotNet
{
	class Program
	{
		// index of the greatest common divisor of each ball
		private static int[] gcdIndex;
		private static int gcdNeeded;


		static void Main(string[] args)
		{
			int size = 0;

			while (int.TryParse(Console.ReadLine(), out size))
			{
				if (size == 0) break;
				if (size >= 27 && size <= 127)
				{
					gcdIndex = new int[size];
					gcdNeeded = size;
					Clock target = new Clock(size, CheckIndex);
					while (gcdNeeded > 0)
						target.Increment();
					PrintResult(size);
				}
			}
		}

		private static void PrintResult(int size)
		{
			var result = 1;
			var lcdSet = gcdIndex.Distinct().ToArray();

			foreach (var lcd in lcdSet)
				result = LCM(result, lcd);

			Console.WriteLine("{0} balls cycle after {1} days.", size, result / 2);
		}

		public static int LCM(int a, int b)
		{
			int max = Math.Max(a, b);
			int min = Math.Min(a, b);

			for (int i = 1; i <= min; i++)
				if ((max * i) % min == 0)
					return i * max;

			return min;
		}

		private static void CheckIndex(int[] order, Clock clock)
		{
			if (gcdNeeded == 0) return;

			for (int i = 0; i < order.Length; i++)
			{
				if (order[i] == i + 1 && gcdIndex[i] == 0)
				{
					gcdNeeded--;
					gcdIndex[i] = clock.Cycles;
				}
			}
		}
	}
}
