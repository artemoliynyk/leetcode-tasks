using Xunit;
using Task1;

namespace Task1.Test {

public class UnitTest1
{
    [Fact]
    public void TwoSumTest()
    {
        var solution = new Solution();
        var result = solution.TwoSum(new int[]{1,2,3}, 5);
        Assert.True(false);
    }
}
}