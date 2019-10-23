using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using System.IO;
using System.Net;

namespace package
{
    public partial class FormMain : Form
    {
        public class T_PackInfo
        {
            public int Number { get; set; }
            public string Name { get; set; }
            public string Path { get; set; }
        }

        private List<T_PackInfo> m_vecPackInfo = new List<T_PackInfo>();

        public FormMain()
        {
            InitializeComponent();
        }

        private void FormMain_Load(object sender, EventArgs e)
        {
            // FormMain Initial
            AutoScaleMode = AutoScaleMode.Dpi;
            MaximizeBox = false;
            MinimizeBox = true;
            DoubleBuffered = true;
            StartPosition = FormStartPosition.CenterScreen;

            // TabPack Initial
            listView_pack.View = View.Details;
            listView_pack.CheckBoxes = true;
            listView_pack.Columns.Add("number", 60, HorizontalAlignment.Center);
            listView_pack.Columns.Add("files", 240, HorizontalAlignment.Center);
            listView_pack.Columns.Add("detials", 400, HorizontalAlignment.Center);

            comboBox_pack.DropDownStyle = ComboBoxStyle.DropDownList;
            comboBox_pack.Items.Clear();
            comboBox_pack.Items.Add("aes");
            comboBox_pack.Items.Add("3des");
            comboBox_pack.Items.Add("des");
            comboBox_pack.Items.Add("rsa");
            comboBox_pack.Items.Add("base64");
            comboBox_pack.SelectedIndex = 0;

            textBox_pack_path.ReadOnly = true;

        }

        private void FormMain_FormClosed(object sender, FormClosedEventArgs e)
        {

        }

        private void Update_pack_listView()
        {
            listView_pack.Items.Clear();
            listView_pack.BeginUpdate();
            for (int i = 0; i < m_vecPackInfo.Count; ++i)
            {
                ListViewItem item = new ListViewItem();
                item.Text = m_vecPackInfo[i].Number.ToString();
                item.SubItems.Add(m_vecPackInfo[i].Name);
                item.SubItems.Add(m_vecPackInfo[i].Path);
                listView_pack.Items.Add(item);
            }
            listView_pack.EndUpdate();
        }

        private void Button_pack_add_Click(object sender, EventArgs e)
        {
            OpenFileDialog dialog = new OpenFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.Multiselect = true;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                foreach (var i in dialog.FileNames)
                {
                    bool flag = false;
                    foreach (var j in m_vecPackInfo)
                    {
                        if(j.Path == i)
                        {
                            flag = true;
                            break;
                        }
                    }
                    if(!flag)
                    {
                        T_PackInfo packInfo = new T_PackInfo();
                        FileInfo file = new FileInfo(i);
                        packInfo.Number = m_vecPackInfo.Count + 1;
                        packInfo.Name = file.Name;
                        packInfo.Path = i;
                        m_vecPackInfo.Add(packInfo);
                    }
                }

                Update_pack_listView();
            }

        }

        private void Button_pack_delete_Click(object sender, EventArgs e)
        {
            if(listView_pack.CheckedItems.Count > 0)
            {
                for (int i = 0; i < listView_pack.CheckedItems.Count; ++i)
                {
                    for (int j = 0; j < m_vecPackInfo.Count;)
                    {
                        if(m_vecPackInfo[j].Number.ToString() == listView_pack.CheckedItems[i].Text)
                        {
                            m_vecPackInfo.RemoveAt(j);
                        }
                        else
                        {
                            ++j;
                        }
                    }
                }

                for(int i = 0; i < m_vecPackInfo.Count; ++i)
                {
                    m_vecPackInfo[i].Number = i + 1;
                }

                Update_pack_listView();
            }
        }

        private void Button_pack_select_Click(object sender, EventArgs e)
        {
            // check pack file name
            if(textBox_pack_file.Text == "")
            {
                MessageBox.Show("Please enter pack file name!", "Warning", MessageBoxButtons.OK, MessageBoxIcon.Warning);
                return;
            }

            // check ext name
            if(!textBox_pack_file.Text.Contains("."))
            {
                if(DialogResult.No == MessageBox.Show("It's show no ext name, do you want to continue?", "Warning", MessageBoxButtons.YesNo, MessageBoxIcon.Question))
                {
                    return;
                }
            }

            SaveFileDialog dialog = new SaveFileDialog();
            dialog.Filter = "All files(*.*)|*.*";
            dialog.FileName = textBox_pack_file.Text;
            dialog.RestoreDirectory = true;

            DialogResult dr = dialog.ShowDialog();
            if (dr == DialogResult.OK)
            {
                textBox_pack_path.Text = dialog.FileName;
            }

        }

        private void Button_pack_execute_Click(object sender, EventArgs e)
        {
            string src = "";
            string dest = "";
            string type = "";
            for(int i = 0; i < m_vecPackInfo.Count - 1; ++i)
            {
                src += string.Format("\"{0}\",", m_vecPackInfo[i].Path);
            }
            src += string.Format("\"{0}\"", m_vecPackInfo[m_vecPackInfo.Count - 1].Path);
            dest = string.Format("\"{0}\"", textBox_pack_path.Text);
            switch(comboBox_pack.SelectedIndex)
            {
                case 0:
                    type = "\"aes\"";
                    break;
                case 1:
                    type = "\"3des\"";
                    break;
                case 2:
                    type = "\"des\"";
                    break;
                case 3:
                    type = "\"rsa\"";
                    break;
                case 4:
                    type = "\"base64\"";
                    break;
                default:
                    type = "\"aes\"";
                    break;
            }

            string body = "";
            body += "{";
            body += string.Format("\"src\":[{0}],\"dest\":{1},\"type\":{2}", src, dest, type);
            body += "}";
            body = body.Replace('\\', '/');

            byte[] requestBody = Encoding.ASCII.GetBytes(body);

            HttpWebRequest request = (HttpWebRequest)WebRequest.Create("http://localhost:8080/satellite/pack");
            request.Method = "POST";
            request.ContentType = "application/json";
            request.ContentLength = requestBody.Length;
            using (Stream reqStream = request.GetRequestStream())
            {
                reqStream.Write(requestBody, 0, requestBody.Length);
            }
            using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
            {
                
            }
        }
    }
}
