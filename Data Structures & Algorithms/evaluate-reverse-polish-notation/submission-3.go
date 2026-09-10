func evalRPN(tokens []string) int {
    operations := []int{}
    var head int
    for _, val := range tokens {
        head = len(operations)
        switch val {
            case "+":
            op1 := operations[head - 1]
            op2 := operations[head - 2]
            operations[head - 2] = op2 + op1
            operations = operations[:head - 1]
        case "-": 
            op1 := operations[head - 1]
            op2 := operations[head - 2]
            operations[head - 2] = op2 - op1
            operations = operations[:head - 1]
        case "*":
            op1 := operations[head - 1]
            op2 := operations[head - 2]
            operations[head - 2] = op2 * op1
            operations = operations[:head - 1]
        case "/":
            op1 := operations[head - 1]
            op2 := operations[head - 2]
            operations[head - 2] = op2 / op1
            operations = operations[:head - 1]
        default:
            num, _ := strconv.Atoi(val)
            operations = append(operations, num)
        }
    }
    
    return operations[0]
}
